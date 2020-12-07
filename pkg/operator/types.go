package operator

import (
	"context"

	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	jinghzhuv1 "github.com/jinghzhu/KubernetesCRD/pkg/crd/jinghzhu/v1"
	jinghzhuv1clientset "github.com/jinghzhu/KubernetesCRD/pkg/crd/jinghzhu/v1/apis/clientset/versioned"
)

const (
	maxRetries int = 3 // Max retry number in the worker queue.
)

// Operator for controlling CRD instances.
type Operator struct {
	// kubeClientset is a standard Kubernetes clientset.
	kubeClientset kubernetes.Interface
	// jinghzhuV1Clientset is a clientset for the sample CRD, which is Jinghzhu v1.
	jinghzhuV1Clientset jinghzhuv1clientset.Interface
	// queue is a rate limited work queue. This is used to queue work to be processed instead of
	// performing it as soon as a change happens. This means we can ensure we only process a fixed
	// amount of resources at a time, and makes it easy to ensure we are never processing the same
	// item simultaneously in two different workers.
	queue        workqueue.RateLimitingInterface
	informer     cache.SharedIndexInformer
	context      context.Context
	crdNamespace string
	podNamespace string
}

// New creates the CRD Operator. The parameter nsOp is the namespace where this Operator will run.
func New(nsOp, nsCRD, nsPod string, kubeClient kubernetes.Interface, jinghzhuV1Client jinghzhuv1clientset.Interface) *Operator {
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	lw := cache.NewListWatchFromClient(jinghzhuV1Client.JinghzhuV1().RESTClient(), jinghzhuv1.Plural, nsOp, fields.Everything())
	// Use SharedIndexInformer instead of SharedInformer because it allows Operator to maintain indexes
	// across all objects in the cache.
	informer := cache.NewSharedIndexInformer(
		lw,
		&jinghzhuv1.Jinghzhu{},
		0, //Skip resync
		cache.Indexers{},
	)
	id, _ := uuid.NewRandom()
	ctx := context.WithValue(context.Background(), "sample-id", id)
	c := &Operator{
		context:             ctx,
		kubeClientset:       kubeClient,
		jinghzhuV1Clientset: jinghzhuV1Client,
		queue:               queue,
		informer:            informer,
		crdNamespace:        nsCRD,
		podNamespace:        nsPod,
	}
	// Events in the Workqueue are represented by their keys which are constructed in the format of
	// crd_instance_namespace/crd_instance_name. In the case of Pod deletion, must check for the DeletedFinalStateUnknown
	// state of that Jinghzhu instance in the cache before enqueuing its key. The DeletedFinalStateUnknown state
	// means that the Pod has been deleted but that the watch deletion event was missed and the Operator didn't
	// react accordingly.
	c.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onAdd,
		UpdateFunc: c.onUpdate,
		DeleteFunc: c.onDelete,
	})

	return c
}

// GetCRDNamespace returns the namespace where Operator watchs the CRs.
func (c *Operator) GetCRDNamespace() string {
	return c.crdNamespace
}

// GetPodNamespace returns the namespace where Operator process the worker Pods.
func (c *Operator) GetPodNamespace() string {
	return c.podNamespace
}

// HasSynced is required for the cache.Controller interface.
func (c *Operator) HasSynced() bool {
	return c.informer.HasSynced()
}

// GetContext retruns the context.
func (c *Operator) GetContext() context.Context {
	return c.context
}
