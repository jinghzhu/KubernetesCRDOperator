package pod

import (
	"strings"

	corev1 "k8s.io/api/core/v1"
)

// GetPodImages returns all container images for the given Pod in a string array.
func GetPodImages(pod *corev1.Pod) []string {
	if pod == nil || pod.Spec.Containers == nil {
		return make([]string, 0)
	}

	images := make([]string, len(pod.Spec.Containers))
	for k, v := range pod.Spec.Containers {
		images[k] = v.Image
	}

	return images
}

// GetPodCommands returns all container commands for the given Pod in a string array.
func GetPodCommands(pod *corev1.Pod) []string {
	if pod == nil || pod.Spec.Containers == nil {
		return make([]string, 0)
	}

	cmds := make([]string, len(pod.Spec.Containers))
	for k, v := range pod.Spec.Containers {
		cmds[k] = strings.Join(v.Command, " ")
	}

	return cmds
}

// IsCompleted retruns true if for each container of the Pod, its State is Terminated and Reason is Completed.
// Although Pod is in Succeeded status, it sometimes doesn't mean all containers were run as expected.
// For example, container may be end by OOMKilled but its Status is Succeeded and the container's State is Terminated.
func IsCompleted(pod *corev1.Pod) bool {
	if pod == nil || pod.Status.Phase != corev1.PodSucceeded || pod.Status.ContainerStatuses == nil {
		return false
	}
	if pod.Spec.Containers == nil || len(pod.Spec.Containers) == 0 {
		return true
	}
	for _, v := range pod.Status.ContainerStatuses {
		if v.State.Terminated == nil || v.State.Terminated.Reason != "Completed" {
			return false
		}
	}

	return true
}
