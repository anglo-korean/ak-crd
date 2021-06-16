package main

import (
	"os"
	"os/user"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func config() (config *rest.Config, err error) {
	if _, inCluster := os.LookupEnv("KUBERNETES_SERVICE_HOST"); inCluster {
		return rest.InClusterConfig()
	}

	var (
		kubeconfig string
		ok         bool
	)

	if kubeconfig, ok = os.LookupEnv("KUBE_CONFIG"); ok {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	u, err := user.Current()
	if err != nil {
		return
	}

	home := u.HomeDir

	kubeconfig = filepath.Join(home, ".kube", "config")

	return clientcmd.BuildConfigFromFlags("", kubeconfig)
}
