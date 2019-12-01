package k8s

import (
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)

type Namespace struct {
	Name   string
	Labels map[string]string
}

type Pod struct {
	Name   string
	Labels map[string]string
	Status string
}

type Cluster interface {
	ListNamespaces() (namespaces []Namespace, err error)
	ListPods(namespace string) (pods []Pod, err error)
}

func NewCluster() Cluster {
	if os.Getenv("APP_MODE") == "mock" {
		logrus.Info("running in APP_MODE=mock")
		return ClusterMock{}
	}
	return k8sCluster{client: createK8sClient()}
}

type k8sCluster struct {
	client *kubernetes.Clientset
}

func (k k8sCluster) ListNamespaces() (namespaces []Namespace, err error) {
	var nsList *v1.NamespaceList
	nsList, err = k.client.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return
	}

	namespaces = make([]Namespace, len(nsList.Items))
	for index, ns := range nsList.Items {
		namespaces[index] = mapNamespace(ns)
	}

	return
}

func mapNamespace(namespace v1.Namespace) Namespace {
	return Namespace{
		Name:   namespace.Name,
		Labels: namespace.Labels,
	}
}

func (k k8sCluster) ListPods(namespace string) (pods []Pod, err error) {
	var podList *v1.PodList
	podList, err = k.client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return
	}

	pods = make([]Pod, len(podList.Items))
	for index, pod := range podList.Items {
		pods[index] = mapPod(pod)
	}

	return
}

func mapPod(pod v1.Pod) Pod {
	return Pod{
		Name:   pod.Name,
		Labels: pod.Labels,
		Status: pod.Status.Message,
	}
}

func createK8sClient() *kubernetes.Clientset {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}
