package main

import (
	"encoding/json"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net/http"
)

var clientset *kubernetes.Clientset

func main() {
	createK8sClient()
	http.HandleFunc("/namespaces", listNamespaces)
	http.HandleFunc("/pods", listPods)
	http.ListenAndServe(":8080", nil)
}

func listNamespaces(w http.ResponseWriter, r *http.Request) {
	var nsList *v1.NamespaceList
	nsList, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(nsList)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func listPods(w http.ResponseWriter, r *http.Request) {
	namespace := r.URL.Query()["ns"][0]

	podList, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(podList)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func createK8sClient() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}
