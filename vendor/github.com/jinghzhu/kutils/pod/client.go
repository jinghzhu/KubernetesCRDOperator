package pod

import (
	"bytes"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/jinghzhu/goutils/utils"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
)

// CreatePod creates a Pod.
func (c *Client) CreatePod(spec *corev1.Pod, namespace string, opts metav1.CreateOptions) (*corev1.Pod, error) {
	return c.kubeClient.CoreV1().Pods(namespace).Create(c.GetContext(), spec, opts)
}

// CreatePodWithRetry creates a Pod with retry.
func (c *Client) CreatePodWithRetry(spec *corev1.Pod, namespace string, opts metav1.CreateOptions) (*corev1.Pod, error) {
	var podCreated *corev1.Pod
	_, err := utils.Retry(30*time.Millisecond, 2, func() (bool, error) {
		pod, err1 := c.kubeClient.CoreV1().Pods(namespace).Create(c.GetContext(), spec, opts)
		if err1 == nil {
			podCreated = pod

			return true, nil
		}
		return false, err1
	})

	return podCreated, err
}

// ListPods returns a list of Pods by namespace and list options.
func (c *Client) ListPods(namespace string, opts metav1.ListOptions) (*corev1.PodList, error) {
	return c.kubeClient.CoreV1().Pods(namespace).List(c.GetContext(), opts)
}

// GetPod returns the Pod instance by namespace, Pod name and get options.
func (c *Client) GetPod(namespace, podName string, opts metav1.GetOptions) (*corev1.Pod, error) {
	return c.kubeClient.CoreV1().Pods(namespace).Get(c.GetContext(), podName, opts)
}

// IsExist returns false if the Pod doesn't exist in the specific namespace.
func (c *Client) IsExist(namespace, podName string) (bool, error) {
	_, err := c.GetPod(namespace, podName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

// AddPodLabel adds a label field into Pod. If the key already exists, it'll overwrite it.
func (c *Client) AddPodLabel(pod *corev1.Pod, key string, value string) (*corev1.Pod, error) {
	if pod.GetLabels() == nil {
		pod.SetLabels(make(map[string]string))
	}
	pod.Labels[key] = value

	return c.UpdatePod(pod, pod.GetNamespace(), metav1.UpdateOptions{})
}

// AddAnnotation adds a new key-value pair into Pod annotation field.
func (c *Client) AddAnnotation(pod *corev1.Pod, key, value string) (*corev1.Pod, error) {
	if pod.GetAnnotations() == nil {
		pod.SetAnnotations(make(map[string]string))
	}
	pod.Annotations[key] = value

	return c.UpdatePod(pod, pod.GetNamespace(), metav1.UpdateOptions{})
}

// UpdatePod accepts a context, pod and namespace. It returns a pointer to pod and error.
func (c *Client) UpdatePod(pod *corev1.Pod, namespace string, opts metav1.UpdateOptions) (*corev1.Pod, error) {
	return c.kubeClient.CoreV1().Pods(namespace).Update(c.GetContext(), pod, opts)
}

// DeletePod talks to Kubernetes to delete a Pod by given delete options.
func (c *Client) DeletePod(namespace, podName string, opts metav1.DeleteOptions) error {
	return c.kubeClient.CoreV1().Pods(namespace).Delete(c.GetContext(), podName, opts)
}

// DeletePodWithCheck deletes the Pod and will start a goroutine in background
// to confirm whether the Pod is successfully deleted.
func (c *Client) DeletePodWithCheck(namespace, podName string, opts metav1.DeleteOptions) error {
	if *opts.GracePeriodSeconds < int64(0) {
		opts.GracePeriodSeconds = &defaultDeletePeriod
	}
	err := c.DeletePod(namespace, podName, opts)
	if err != nil {
		return err
	}
	go c.WaitForDeletion(namespace, podName, opts.GracePeriodSeconds)

	return nil
}

// WaitForDeletion will wait for a period to check if Pod is deleted.
func (c *Client) WaitForDeletion(namespace, podName string, period *int64) error {
	var waitTime time.Duration
	if period != nil {
		waitTime = time.Duration(*period) * time.Second
	} else {
		waitTime = time.Duration(defaultDeletePeriod) * time.Second
	}

	time.Sleep(waitTime) // Wait for gracefully deletion.

	// Check if the Pod is deleted. If it sill exits, we'll force to delete it. Then check it again.
	// This logic will be tried 3 times at max.
	err := wait.Poll(
		time.Duration(defaultDeletePeriod)*time.Second,
		3*time.Duration(defaultDeletePeriod)*time.Second,
		func() (bool, error) {
			exist, err := c.IsExist(namespace, podName)
			if err != nil {
				return false, err
			} else if exist {
				var gracePeriod int64
				c.DeletePod(namespace, podName, metav1.DeleteOptions{GracePeriodSeconds: &gracePeriod})

				return false, nil
			}

			return true, nil
		},
	)

	return err
}

// GetLogString returns the log of Pod in string.
func (c *Client) GetLogString(namespace, podName string, opts *corev1.PodLogOptions) (string, error) {
	stream, err := c.kubeClient.CoreV1().Pods(namespace).GetLogs(podName, opts).Stream(c.GetContext())
	defer func() {
		if stream != nil {
			stream.Close()
		}
	}()
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)

	return buf.String(), nil
}

// GetEvents returns a EventList object for the given Pod name and list options.
func (c *Client) GetEvents(namespace, podName string, opts metav1.ListOptions) (*corev1.EventList, error) {
	return c.kubeClient.CoreV1().Events(namespace).List(c.GetContext(), opts)
}

// GetVersion returns the version of the the current REST client
func (c *Client) GetVersion() schema.GroupVersion {
	return c.kubeClient.CoreV1().RESTClient().APIVersion()
}
