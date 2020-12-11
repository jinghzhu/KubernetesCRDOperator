# Kubernetes CRD Controller
This repository is to provide a sample about how to create a controller to watch all events (add/update/delete) of some CRD kinds so that we can implement the Operator pattern.

The CRD kind referred here is [Jinghzhu v1](https://github.com/jinghzhu/KubernetesCRD). Of course, you can use any other. The controller will do something like [ReplicaSet](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) against the CRD.

Assume:
1. Go version > 1.9.0
2. A Kubernetes cluster is available.
3. Kubernetes version >= 1.18.0



# How It Works
There are many articles introducing the mechanism of Operator in Kubernetes community. So, I wouldn't pay too much attention on it. In short, it leverage the etcd watch mechanism to catch all events.

In the method `New()` at `pkg/operator/types.go`, we define the major components of Operator:
1. queue - I implements an in-memory queue to process all events in concurrent way. By default, there is only one main goroutine to help the Operator get events from etcd and call add/update/delete handlers. Now, after receiving events, it puts them in the queue which can ensure the concurrent safe.
2. lister
3. informer
4. event handler callback

```go
queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
lw := cache.NewListWatchFromClient(jinghzhuV1Client.JinghzhuV1().RESTClient(), jinghzhuv1.Plural, namespace, fields.Everything())

...

informer := cache.NewSharedIndexInformer(
	lw,
	&jinghzhuv1.Jinghzhu{},
	0, //Skip resync
	cache.Indexers{},
)

...

c.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
	AddFunc:    c.onAdd,
	UpdateFunc: c.onUpdate,
	DeleteFunc: c.onDelete,
})
```

For details, please see `pkg/operator/operator.go` and `pkg/operator/worker.go`.



# How It Looks Like
```bash
$ go run cmd/operator/main.go

Init CRD Operator...
Ready to start CRD Operator...
Find an onAdd event
        key = crd/jinghzhu-example-f7wgv
        eventType = add
        old = <nil>
        new=    Name = jinghzhu-example-f7wgv
        Resource Version = 2343534
        Desired = 1
        Current = 0
        PodList =
        State = Pending
        Message = Created but not processed yet

CRD Operator is synced and ready...
Ready to start CRD Operator workers...
Processing Jinghzhu instance    key = crd/jinghzhu-example-f7wgv
        eventType = add
        old = <nil>
        new=    Name = jinghzhu-example-f7wgv
        Resource Version = 2343534
        Desired = 1
        Current = 0
        PodList =
        State = Pending
        Message = Created but not processed yet

All workers are started...
Ready to reconcile add event for        Name = jinghzhu-example-f7wgv
        Resource Version = 2343534
        Desired = 1
        Current = 0
        PodList =
        State = Pending
        Message = Created but not processed yet

Finish reconciling add event for        Name = jinghzhu-example-f7wgv
        Resource Version = 2343534
        Desired = 1
        Current = 1
        PodList = jinghzhu-worker-hl7fc
        State = Running
        Message = Start to run replica set

Find an onUpdate event
        key = crd/jinghzhu-example-f7wgv
        eventType = update
        old =   Name = jinghzhu-example-f7wgv
        Resource Version = 2343534
        Desired = 1
        Current = 0
        PodList =
        State = Pending
        Message = Created but not processed yet

        new=    Name = jinghzhu-example-f7wgv
        Resource Version = 2347327
        Desired = 1
        Current = 1
        PodList = jinghzhu-worker-hl7fc
        State = Running
        Message = Start to run replica set

Processing Jinghzhu instance    key = crd/jinghzhu-example-f7wgv
        eventType = update
        old =   Name = jinghzhu-example-f7wgv
        Resource Version = 2343534
        Desired = 1
        Current = 0
        PodList =
        State = Pending
        Message = Created but not processed yet

        new=    Name = jinghzhu-example-f7wgv
        Resource Version = 2347327
        Desired = 1
        Current = 1
        PodList = jinghzhu-worker-hl7fc
        State = Running
        Message = Start to run replica set

Ready to reconcile update event for     Name = jinghzhu-example-f7wgv
        Resource Version = 2347327
        Desired = 1
        Current = 1
        PodList = jinghzhu-worker-hl7fc
        State = Running
        Message = Start to run replica set

No need to reconcile jinghzhu-example-f7wgv
Finish reconciling update event for     Name = jinghzhu-example-f7wgv
        Resource Version = 2347327
        Desired = 1
        Current = 1
        PodList = jinghzhu-worker-hl7fc
        State = Running
        Message = Start to run replica set
```