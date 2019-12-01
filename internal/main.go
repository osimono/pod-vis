package main

import (
	"github.com/osimono/pod-vis/internal/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/namespaces", handler.ListNamespaces)
	http.HandleFunc("/pods", handler.ListPods)
	http.ListenAndServe(":8080", nil)
}
