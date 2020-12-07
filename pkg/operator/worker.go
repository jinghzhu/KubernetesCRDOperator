package operator

import (
	"fmt"

	"github.com/jinghzhu/KubernetesCRDOperator/pkg/client"
	"github.com/jinghzhu/KubernetesCRDOperator/pkg/config"
	"github.com/jinghzhu/KubernetesCRDOperator/pkg/events"
	"github.com/jinghzhu/KubernetesCRDOperator/pkg/types"

	jinghzhuv1client "github.com/jinghzhu/KubernetesCRD/pkg/crd/jinghzhu/v1/client"
	crdtypes "github.com/jinghzhu/KubernetesCRD/pkg/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func (c *Operator) runWorker() {
	// It will automatically wait until there's some work item available.
	for c.processNextItem() {
	}
}

// processNextWorkItem deals with one key off the queue. Return false when it's time to quit.
func (c *Operator) processNextItem() bool {
	// Wait for an event or signal to quit.
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
		fmt.Printf("Ready to reconcile %s event for %s\n", newEvent.EventType, newEvent.NewJinghzhu.String())
		err = reconcile(newEvent)
		fmt.Printf("Finish reconciling %s event for %s\n", newEvent.EventType, newEvent.NewJinghzhu.String())

		return err
	case events.EventUpdate:
		fmt.Printf("Ready to reconcile %s event for %s\n", newEvent.EventType, newEvent.NewJinghzhu.String())
		err = reconcile(newEvent)
		fmt.Printf("Finish reconciling %s event for %s\n", newEvent.EventType, newEvent.NewJinghzhu.String())

		return err
	case events.EventDelete:
		fmt.Printf("Ready to reconcile %s event for %s\n", newEvent.EventType, newEvent.NewJinghzhu.String())
		err = reconcile(newEvent)
		fmt.Printf("Finish reconciling %s event for %s\n", newEvent.EventType, newEvent.NewJinghzhu.String())

		return err
	default:
		fmt.Printf("No case match Jinghzhu instance %s\n", newEvent.String())
	}

	return nil
}

func reconcile(event events.Event) error {
	new := event.NewJinghzhu
	cfg := config.GetConfig()
	crdName, crdNamespace, podNamespace := new.GetName(), new.GetNamespace(), cfg.GetPodNamespace()

	if new.Spec.Desired == new.Spec.Current {
		fmt.Printf("No need to reconcile %s\n", crdName)

		return nil
	}

	podClient := client.GetDefaultPodClient()
	crdClient, err := jinghzhuv1client.NewClient(types.GetDefaultCtx(), cfg.GetKubeconfigPath(), crdNamespace)
	if err != nil {
		return err
	}
	defaultPodSpec, err := types.GetDefaultPodSpec()
	if err != nil {
		fmt.Printf("Fail to get default Pod spec in reconcile: %+v\n", err)

		return err
	}
	defaultPodSpec.Labels["crd"] = crdName

	if new.Spec.Desired > new.Spec.Current {
		delta := new.Spec.Desired - new.Spec.Current

		for i := 0; i < delta; i++ {
			pod, err := podClient.CreatePodWithRetry(defaultPodSpec, podNamespace, metav1.CreateOptions{})
			if err != nil {
				fmt.Printf("Fail to create Pod in reconcile: %+v\n", err)

				return err
			}
			new.Spec.PodList = append(new.Spec.PodList, pod.GetName())
		}
		if new.Status.State == crdtypes.StatePending {
			new.Status.State = crdtypes.StateRunning
			new.Status.Message = "Start to run replica set"
		} else {
			new.Status.Message = "Scaling Out"
		}
		new.Spec.Current = new.Spec.Desired
	}

	if new.Spec.Desired < new.Spec.Current {
		delta := new.Spec.Current - new.Spec.Desired
		for i := 0; i < delta; i++ {
			err = podClient.DeletePod(podNamespace, new.Spec.PodList[i], metav1.DeleteOptions{})
			if err != nil {
				fmt.Printf("Fail to delete Pod %s in reconcile: %+v\n", new.Spec.PodList[i], err)

				return err
			}
		}
		new.Spec.PodList = new.Spec.PodList[delta:]
		new.Status.Message = "Scaling In"
		new.Spec.Current = new.Spec.Desired
	}

	_, err = crdClient.PatchSpecAndStatus(crdName, &new.Spec, &new.Status)
	if err != nil {
		fmt.Printf("Fail to patch CRD status and spec in reconcile: %+v\n", err)
	}

	return err
}
