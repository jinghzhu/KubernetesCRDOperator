package config

var (
	config *Config
)

const (
	// DefaultPodNamespace is the default namespace where the CRD Operator will create worker Pods.
	DefaultPodNamespace string = "worker"
	// DefaultCRDNamespace is the default namespace where the CRD Operator will process CRs.
	DefaultCRDNamespace string = "crd"
	// DefaultKubeconfigPath is the default local path of kubeconfig file.
	DefaultKubeconfigPath string = "/.kube/config"
)

type Config struct {
	podNamespace   string
	crdNamespace   string
	kubeconfigPath string
}

func (c *Config) GetPodNamespace() string {
	return c.podNamespace
}

func (c *Config) GetCRDNamespace() string {
	return c.crdNamespace
}

func (c *Config) GetKubeconfigPath() string {
	return c.kubeconfigPath
}
