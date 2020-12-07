package config

import "os"

func init() {
	initConfig()
}

func initConfig() {
	config = &Config{}
	config.podNamespace = os.Getenv("CO_POD_NAMESPACE")
	if config.podNamespace == "" {
		config.podNamespace = DefaultPodNamespace
	}
	config.crdNamespace = os.Getenv("CO_CRD_DNAMESPACE")
	if config.crdNamespace == "" {
		config.crdNamespace = DefaultCRDNamespace
	}
	config.kubeconfigPath = os.Getenv("CO_KUBECONFIG")
	if config.kubeconfigPath == "" {
		config.kubeconfigPath = DefaultKubeconfigPath
	}
}

func GetConfig() *Config {
	return config
}
