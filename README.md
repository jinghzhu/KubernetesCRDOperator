# Kubernetes CRD Operator
This repository is to provide a sample about how to create an Operator to watch all events (add/update/delete) of some CRD kinds.

The CRD kind refered here is [Jinghzhu v1](https://github.com/jinghzhu/KubernetesCRD). Of course, you can use any other.

Assume:
1. Go version > 1.9.0
2. A Kubernetes cluster is available.
3. Kubernetes version >= 1.9.6



# How It Work
There are many articles introducing the mechanism of Operator in Kubernetes community. So, I wouldn't pay too much attention on it. In short, it leverage the etcd watch mechanism to catch all events.

In the method `New()` at `pkg/operator/types.go`, we define what we are interested in and corresponding handlers:
```go
cache.NewListWatchFromClient(jinghzhuV1Client.JinghzhuV1().RESTClient(), jinghzhuv1.Plural, namespace, fields.Everything())
```

```go
c.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
	AddFunc:    c.onAdd,
	UpdateFunc: c.onUpdate,
	DeleteFunc: c.onDelete,
})
```

Different with many other Operators, mine implements a in-memory queue to process all events in concurrent way. By default, there is only one main goroutine to help the Operator get events from etcd and call add/update/delete handlders. Now, after receving events, it puts them in the queue which can ensure the concurrent safe.

For details, please see `pkg/operator/operator.go` and `pkg/operator/worker.go`.



# How It Looks Like
```bash
$ KUBECONFIG=~/.kube/config go run cmd/operator/main.go

Init CRD Operator...

Ready to start CRD Operator

CRD Operator is synced and ready

Ready to start CRD Operator workers...

All workers are started

Processing Jinghzhu instance Event: key = crd-ns/jinghzhu-example1, eventType = add, namespace = , oldJinghzhu = <nil>, newJinghzhu= &{TypeMeta:{Kind:Jinghzhu APIVersion:jinghzhu.com/v1} ObjectMeta:{Name:jinghzhu-example1 GenerateName: Namespace:crd-ns SelfLink:/apis/jinghzhu.com/v1/namespaces/crd-ns/jinghzhus/jinghzhu-example1 UID:2095f6e4-8b76-11e9-92f8-02005ee2a828 ResourceVersion:40895595 Generation:0 CreationTimestamp:2019-06-10 19:51:48 +0800 CST DeletionTimestamp:<nil> DeletionGracePeriodSeconds:<nil> Labels:map[] Annotations:map[] OwnerReferences:[] Initializers:nil Finalizers:[] ClusterName:} Spec:{Foo:hello Bar:true} Status:{State:Pending Message:Created but not processed yet}}

Processing Jinghzhu instance jinghzhu-example1 in created case
```