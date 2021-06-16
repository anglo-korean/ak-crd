package main

import (
	"fmt"
	"time"

	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/klog/v2"

	// Our versioned types
	"anglo-korean.github.io/crd/v1alpha1"
)

var (
	// Set during build
	version string

	// These are a list of versions we care about, both for creating CRDs
	// and for running informers/ watchers. We can promote versions in and
	// out of this when we need to.
	publishedVersions = []string{"v1alpha1"}

	// Various CRD creation/ registration functions we've created
	crdCreaters = map[string]CrdCreater{
		"v1alpha1": v1alpha1.CreateCRD,
	}
)

type CrdCreater func(apiextension.Interface) error

func createCrds(clientset apiextension.Interface) (err error) {
	for _, v := range publishedVersions {
		err = crdCreaters[v](clientset)
		if err != nil {
			return fmt.Errorf("%s -> %w", v, err)
		}
	}

	return nil
}

func init() {
	klog.InitFlags(nil)
}

func main() {
	config, err := config()
	if err != nil {
		klog.Fatalf("error finding kubernetes: %v", err)
	}

	kubeClient, err := apiextension.NewForConfig(config)
	if err != nil {
		klog.Fatalln("error creating kubernetes connection: %v", err)
	}

	// CreateCRDs
	err = createCrds(kubeClient)
	if err != nil {
		klog.Fatalf("error creating CRDs: %v", err)
	}

	// Wait for the CRD to be created before we use it.
	time.Sleep(5 * time.Second)
}
