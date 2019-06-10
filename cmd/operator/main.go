package main

import (
	"fmt"
	"os"

	jinghzhuv1clientset "github.com/jinghzhu/KubernetesCRD/pkg/crd/jinghzhu/v1/apis/clientset/versioned"

	"github.com/jinghzhu/KubernetesCRDOperator/pkg/operator"
	"k8s.io/client-go/kubernetes"
)

func main() {
	fmt.Println("Init CRD Operator...")
	// Use kubeconfig to create client config.
	clientConfig, err := config.GetRESTClientConfig("", os.Getenv("KUBECONFIG"))
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
	operator := operator.New(kubeClient, crdClientset)
	
	go operator.Run(2)

	ch := make(chan bool, 1)
	<-ch
}