package main

import (
	"fmt"

	jinghzhuv1clientset "github.com/jinghzhu/KubernetesCRD/pkg/crd/jinghzhu/v1/apis/clientset/versioned"

	"github.com/jinghzhu/KubernetesCRDOperator/pkg/config"
	"github.com/jinghzhu/KubernetesCRDOperator/pkg/operator"
	"github.com/jinghzhu/KubernetesCRDOperator/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	fmt.Println("Init CRD Operator...")
	cfg := config.GetConfig()
	// Use kubeconfig to create client config.
	clientConfig, err := clientcmd.BuildConfigFromFlags("", cfg.GetKubeconfigPath())
	if err != nil {
		panic(err)
	}
	kubeClient, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		panic(err)
	}
	crdClientset, err := jinghzhuv1clientset.NewForConfig(clientConfig)
	if err != nil {
		panic(err)
	}
	nsCRD := cfg.GetCRDNamespace()
	operator := operator.New(nsCRD, nsCRD, cfg.GetPodNamespace(), kubeClient, crdClientset)

	go operator.Run(types.WorkerNum)

	ch := make(chan bool, 1)
	<-ch
}
