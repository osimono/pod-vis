package handler

import (
	"encoding/json"
	"github.com/osimono/pod-vis/internal/k8s"
	"github.com/sirupsen/logrus"
	"net/http"
)

var cluster = k8s.NewCluster()

func ListNamespaces(w http.ResponseWriter, r *http.Request) {
	namespaces, err := cluster.ListNamespaces()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	writeJson(namespaces, w)
}

func ListPods(w http.ResponseWriter, r *http.Request) {
	namespace := r.URL.Query()["ns"][0]

	pods, err := cluster.ListPods(namespace)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	writeJson(pods, w)
}

func writeJson(v interface{}, w http.ResponseWriter) {
	bytes, err := json.Marshal(v)
	if err != nil {
		logrus.Errorf(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
