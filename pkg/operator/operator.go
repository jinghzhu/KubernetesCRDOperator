package operator

import (
	"fmt"

	jinghzhuv1 "github.com/jinghzhu/KubernetesCRD/pkg/crd/jinghzhu/v1"
	events "github.com/jinghzhu/KubernetesCRDOperator/pkg/events"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"errors"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
)

// Run is the main entry to start CRD Operator.
func (c *Operator) Run(workerNum int) error {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown() // Make sure work queue is shutdown which will trigger workers to end.

	fmt.Println("Ready to start CRD Operator")
	go c.informer.Run(c.GetContext().Done())

	// Wait for the caches to synchronize before starting workers.
	if !cache.WaitForCacheSync(c.GetContext().Done(), c.HasSynced) {
		errMsg := "Timed out waiting for cache to sync"
		fmt.Println(errMsg)
		err := errors.New(errMsg)
		utilruntime.HandleError(err)

		return err
	}
	fmt.Println("CRD Operator is synced and ready")

	fmt.Println("Ready to start CRD Operator workers...")
	for i := 0; i < workerNum; i++ {
		// runWorker will loop until some error happens. wait.Until will rekick the worker after one second.
		go wait.Until(c.runWorker, time.Second, c.GetContext().Done())
	}
	fmt.Println("All workers are started")
	<-c.GetContext().Done()
	fmt.Println("Shutting down all workers")

	return c.GetContext().Err()
}

func (c *Operator) onAdd(obj interface{}) {
	instance := obj.(*jinghzhuv1.Jinghzhu).DeepCopy()
	var newEvent events.Event
	var err error
	newEvent.Key, err = cache.MetaNamespaceKeyFunc(obj)
	newEvent.EventType = events.EventAdd
	newEvent.NewJinghzhu = instance
	if err == nil {
		c.queue.Add(newEvent)
	} else {
		fmt.Printf("Error in getting event in onAdd: %+v\n", err)
	}
}

func (c *Operator) onUpdate(oldObj, newObj interface{}) {
	old := oldObj.(*jinghzhuv1.Jinghzhu).DeepCopy()
	new := newObj.(*jinghzhuv1.Jinghzhu).DeepCopy()
	var newEvent events.Event
	var err error
	newEvent.Key, err = cache.MetaNamespaceKeyFunc(newObj)
	newEvent.EventType = events.EventUpdate
	newEvent.OldJinghzhu = old
	newEvent.NewJinghzhu = new
	if err == nil {
		c.queue.Add(newEvent)
	} else {
		fmt.Printf("Error in getting event in onUpdate: %+v\n", err)
	}
}

func (c *Operator) onDelete(obj interface{}) {
	instance := obj.(*jinghzhuv1.Jinghzhu).DeepCopy()
	var newEvent events.Event
	var err error
	newEvent.Key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	newEvent.EventType = events.EventDelete
	newEvent.NewJinghzhu = instance
	if err == nil {
		c.queue.Add(newEvent)
	} else {
		fmt.Printf("Error in getting event in onDelete: %+v\n", err)
	}
}
