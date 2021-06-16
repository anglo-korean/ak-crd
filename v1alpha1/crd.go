package v1alpha1

import (
	"context"
	"reflect"

	apiextensionv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apiextensionv1client "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	CRDGroup   string = "stable.anglo-korean.github.io"
	CRDVersion string = "v1alpha1"

	CronjobPlural string = "cronjobs"
	FQCronjobName string = CronjobPlural + "." + CRDGroup
)

func CreateCRD(clientset apiextension.Interface) error {
	return createCRD(clientset.ApiextensionsV1().CustomResourceDefinitions(), CronjobPlural, FQCronjobName, CronjobConfig{})
}

func createCRD(c apiextensionv1client.CustomResourceDefinitionInterface, plural, fqName string, i interface{}) error {
	crd := &apiextensionv1.CustomResourceDefinition{
		ObjectMeta: meta_v1.ObjectMeta{Name: fqName},
		Spec: apiextensionv1.CustomResourceDefinitionSpec{
			Group:    CRDGroup,
			Versions: []apiextensionv1.CustomResourceDefinitionVersion{{Name: CRDVersion, Storage: true}},
			Scope:    apiextensionv1.NamespaceScoped,
			Names: apiextensionv1.CustomResourceDefinitionNames{
				Plural: plural,
				Kind:   reflect.TypeOf(i).Name(),
			},
		},
	}

	_, err := c.Create(context.TODO(), crd, meta_v1.CreateOptions{})
	if err != nil && apierrors.IsAlreadyExists(err) {
		return nil
	}

	return err
}
