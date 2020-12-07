package operator

import (
	"fmt"

	"github.com/jinghzhu/KubernetesCRDOperator/pkg/events"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func (c *Operator) runWorker() {
	// It will automatically wait until there's some work item available.
	for c.processNextItem() {
	}
}

// processNextWorkItem deals with one key off the queue. Return false when it's time to quit.
func (c *Operator) processNextItem() bool {
	// Wait until there is a new item in the working queue.
	newEvent, quit := c.queue.Get()
	if quit {
		fmt.Println("Return false in processNextItem")

		return false
	}
	// Always indicate to queue that we have completed work. This frees the key for other workers.
	// It provides safe parallel processing because two Pods with the same key are never processed in
	// parallel.
	defer c.queue.Done(newEvent)
	err := c.processItem(newEvent.(events.Event))
	if err == nil {
		// In short, when there is no error, we will reset the ratelimit counters which means tell the queue
		// to stop tracking history.
		// In details, we forget about AddRateLimited history of the key on every successful synchronization.
		// This ensures processing of updates for this key in future is not delayed because of an outdated
		// error history.
		c.queue.Forget(newEvent)
	} else if c.queue.NumRequeues(newEvent) < maxRetries {
		// Let Operator retries `maxRetries` times if error. After that, it stops trying.
		fmt.Printf("Error in processing events and will retry: %+v\n", err)
		// Re-enqueue the key rate limited. Based on the rate limiter on the queue and the re-enqueue history,
		// the key will be processed later again.
		c.queue.AddRateLimited(newEvent)
	} else {
		fmt.Printf("Error and too many retries in processing events: %+v\n", err)
		c.queue.Forget(newEvent) // Remove it if too many error retries.
		utilruntime.HandleError(err)
	}

	return true
}

func (c *Operator) processItem(newEvent events.Event) error {
	_, exists, err := c.informer.GetIndexer().GetByKey(newEvent.Key)
	if err != nil {
		fmt.Printf("Fail to fetch object in processItem: %+v\n", err)

		return err
	}
	// It's an update when Function API object is actually deleted, no need to process anything here.
	// Some words for the delete event:
	//     It means the object is removed in the informer. What is lucky is we can still find its copied
	//     one in the queue. Here, we have already passed everthing of the event, including the object itself.
	//     Another option is we can return a not found error here. In the caller, it is sure it has everything
	//     in the queue. Then we can let caller process it.
	if !exists && newEvent.EventType != events.EventDelete {
		fmt.Println("Jinghzhu instance not found in cache and ignore update")

		return nil
	}
	fmt.Printf("Processing Jinghzhu instance %s\n", newEvent.String())

	// Here to add event based action in future if needed.
	switch newEvent.EventType {
	case events.EventAdd:
		instanceName := newEvent.NewJinghzhu.GetName()
		fmt.Printf("Ready to process Jinghzhu instance %s in created case\n", instanceName)
		fmt.Printf("Finish processing Jinghzhu instance %s in created case\n", instanceName)

		return nil
	case events.EventUpdate:
		old, new := newEvent.OldJinghzhu, newEvent.NewJinghzhu
		instanceName := new.GetName()
		// Only care state change event.
		if old.Status.State != new.Status.State {
			fmt.Printf(
				"Ready to process Jinghzhu instance %s in update case from %s to %s\n",
				instanceName,
				old.Status.State,
				new.Status.State,
			)
			fmt.Printf(
				"Finish processing Jinghzhu instance %s in update case from %s to %s\n",
				instanceName,
				old.Status.State,
				new.Status.State,
			)
		}

		return nil
	case events.EventDelete:
		fmt.Printf("Ready to process Jinghzhu instance %s in delete case\n", newEvent.Key)
		fmt.Printf("Finish processing Jinghzhu instance %s in delete case\n", newEvent.Key)

		return nil
	default:
		fmt.Printf("No case match Jinghzhu instance %s\n", newEvent.String())
	}

	return nil
}
