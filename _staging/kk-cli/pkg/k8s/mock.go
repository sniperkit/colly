package k8s

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/manveru/faker"

	"k8s.io/client-go/kubernetes/fake"

	"github.com/nii236/kk/pkg/kk"

	appsv1 "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockClientSet contains an embedded Kubernetes client set
type MockClientSet struct {
	faker     *faker.Faker
	clientSet *fake.Clientset
}

// NewMock returns a new clientset
func NewMock(flags *k.ParsedFlags) (*MockClientSet, error) {
	mockClientSet := fake.NewSimpleClientset()
	mocker, err := faker.New("en")
	if err != nil {
		return nil, err
	}

	cs := &MockClientSet{
		faker:     mocker,
		clientSet: mockClientSet,
	}

	cs.seed()

	return cs, nil
}

func (cs *MockClientSet) seedNamespaces() error {
	for i := 0; i < 5; i++ {
		_, err := cs.clientSet.CoreV1().Namespaces().Create(&corev1.Namespace{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Namespace",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: strings.Join(cs.faker.Words(2, true), "-"),
				Labels: map[string]string{
					"tag": "mockdata",
				},
				CreationTimestamp: metav1.Time{
					Time: time.Now().Add(-48 * time.Hour),
				},
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (cs *MockClientSet) seedPod(ns, name string) error {
	_, err := cs.clientSet.CoreV1().Pods(ns).Create(&corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"tag": "mockdata",
			},
			CreationTimestamp: metav1.Time{
				Time: time.Now().Add(-48 * time.Hour),
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  cs.faker.Name(),
					Image: cs.faker.Name(),
				},
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// GetDeployments returns all deployments from the k8s cluster
func (cs *MockClientSet) GetDeployments(namespace string) (*appsv1.DeploymentList, error) {
	return cs.clientSet.AppsV1beta1().Deployments(namespace).List(metav1.ListOptions{})
}

func (cs *MockClientSet) seedPods() error {
	namespaces, err := cs.clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, ns := range namespaces.Items {
		for i := 0; i < 5; i++ {
			cs.seedPod(ns.Name, strings.Join(cs.faker.Words(3, true), "-"))
		}
	}
	// os.Exit(1)
	return nil
}

func (cs *MockClientSet) seedDeployments() error {
	for i := 0; i < 5; i++ {
		_, err := cs.clientSet.AppsV1beta1().Deployments("default").Create(&appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Namespace",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: cs.faker.Name(),
				Labels: map[string]string{
					"tag": "mockdata",
				},
				CreationTimestamp: metav1.Time{
					Time: time.Now().Add(-48 * time.Hour),
				},
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
func (cs *MockClientSet) seed() error {
	err := cs.seedNamespaces()
	if err != nil {
		return err
	}
	err = cs.seedDeployments()
	if err != nil {
		return err
	}
	err = cs.seedPods()
	if err != nil {
		return err
	}
	return nil
}

// GetPods (use namespace)
func (cs *MockClientSet) GetPods(namespace string) (*corev1.PodList, error) {
	return cs.clientSet.CoreV1().Pods(namespace).List(metav1.ListOptions{})
}

// GetNamespaces will GetNamespaces
func (cs *MockClientSet) GetNamespaces() (*corev1.NamespaceList, error) {
	return cs.clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})
}

// GetPodContainers will GetPodContainers
func (cs *MockClientSet) GetPodContainers(podName string, namespace string) ([]string, error) {
	result := []string{}
	for i := 0; i < 3; i++ {
		result = append(result, strings.Join(cs.faker.Words(2, true), "-"))
	}
	return result, nil
}

// DeletePod will DeletePod
func (cs *MockClientSet) DeletePod(podName string, namespace string) error {
	return cs.clientSet.CoreV1().Pods(namespace).Delete(podName, &metav1.DeleteOptions{})
}

// GetPodContainerLogs will GetPodContainerLogs
func (cs *MockClientSet) GetPodContainerLogs(podName string, containerName string, namespace string, o io.Writer) error {
	fmt.Fprint(o, strings.Join(cs.faker.Sentences(5, true), "\n"))

	return nil
}
