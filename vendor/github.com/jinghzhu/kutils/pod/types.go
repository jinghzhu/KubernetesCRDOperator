package pod

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	defaultDeletePeriod int64 = 2
)

// Client helps talk to Kubernetes objects.
type Client struct {
	ctx        context.Context
	kubeClient *kubernetes.Clientset
}

// GetContext returns the context of client.
func (c *Client) GetContext() context.Context {
	return c.ctx
}

// New returns a pointer to Client object. If neither masterUrl or kubeconfigPath are passed in we fallback
// to inClusterConfig. If inClusterConfig fails, we fallback to the default config.
func New(ctx context.Context, masterURL, kubeconfigPath string) (*Client, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	clientConfig, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigPath)
	if err != nil {
		return nil, err
	}
	clientSet, err := kubernetes.NewForConfig(clientConfig)

	return &Client{
		ctx:        ctx,
		kubeClient: clientSet,
	}, err
}
