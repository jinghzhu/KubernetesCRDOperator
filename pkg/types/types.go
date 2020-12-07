package types

import (
	"context"

	"github.com/jinghzhu/KubernetesCRDOperator/pkg/config"
	"k8s.io/apimachinery/pkg/api/resource"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	defaultCtx context.Context
)

const (
	// WorkerNum is the number of worker goroutines.
	WorkerNum int = 2
)

func init() {
	defaultCtx = context.Background()
}

func GetDefaultCtx() context.Context {
	return defaultCtx
}

func GetDefaultPodSpec() (*corev1.Pod, error) {
	cfg := config.GetConfig()
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "crd-worker-",
			Namespace:    cfg.GetPodNamespace(),
			Labels:       make(map[string]string),
		},
		Spec: corev1.PodSpec{
			SecurityContext: &corev1.PodSecurityContext{},
			HostNetwork:     false,
			RestartPolicy:   "Never",
		},
	}

	defaultCPU, defaultMem := "300m", "400Mi"
	kubeResource := corev1.ResourceList{}
	cpu, err := resource.ParseQuantity(defaultCPU)
	if err != nil {
		return pod, err
	}
	memory, err := resource.ParseQuantity(defaultMem)
	if err != nil {
		return pod, err
	}
	kubeResource[corev1.ResourceCPU] = cpu
	kubeResource[corev1.ResourceMemory] = memory

	container := corev1.Container{
		Image:           "ubuntu:20.10",
		Command:         []string{"sleep", "60000000"},
		Name:            "crd-worker-container",
		ImagePullPolicy: "IfNotPresent",
		Resources: corev1.ResourceRequirements{
			Limits:   kubeResource,
			Requests: kubeResource,
		},
	}

	pod.Spec.Containers = []corev1.Container{container}

	return pod, nil
}
