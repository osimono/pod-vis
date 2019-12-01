package k8s

type ClusterMock struct {
}

func (c ClusterMock) ListNamespaces() ([]Namespace, error) {
	return []Namespace{
		{Name: "test1", Labels: map[string]string{"Label1": "Value1"}},
		{Name: "some-namespace", Labels: map[string]string{"ddd": "eeeeeee"}},
	}, nil
}

func (c ClusterMock) ListPods(namespace string) ([]Pod, error) {
	return []Pod{
		{
			Name:   "application-pod-123456-kllll",
			Status: "Running",
			Labels: map[string]string{"app": "pod"},
		},
		{
			Name:   "application-pod-123456-kllll",
			Status: "Waiting",
			Labels: map[string]string{"role": "api", "app": "application"},
		},
	}, nil
}
