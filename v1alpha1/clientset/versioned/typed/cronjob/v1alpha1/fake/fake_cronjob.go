/*
Copyright Anglo Korean.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "anglo-korean.github.io/crd/cronjob/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCronjobs implements CronjobInterface
type FakeCronjobs struct {
	Fake *FakeStableV1alpha1
	ns   string
}

var cronjobsResource = schema.GroupVersionResource{Group: "stable.anglo-korean.github.io", Version: "v1alpha1", Resource: "cronjobs"}

var cronjobsKind = schema.GroupVersionKind{Group: "stable.anglo-korean.github.io", Version: "v1alpha1", Kind: "Cronjob"}

// Get takes name of the cronjob, and returns the corresponding cronjob object, and an error if there is any.
func (c *FakeCronjobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Cronjob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cronjobsResource, c.ns, name), &v1alpha1.Cronjob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cronjob), err
}

// List takes label and field selectors, and returns the list of Cronjobs that match those selectors.
func (c *FakeCronjobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CronjobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cronjobsResource, cronjobsKind, c.ns, opts), &v1alpha1.CronjobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CronjobList{ListMeta: obj.(*v1alpha1.CronjobList).ListMeta}
	for _, item := range obj.(*v1alpha1.CronjobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cronjobs.
func (c *FakeCronjobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cronjobsResource, c.ns, opts))

}

// Create takes the representation of a cronjob and creates it.  Returns the server's representation of the cronjob, and an error, if there is any.
func (c *FakeCronjobs) Create(ctx context.Context, cronjob *v1alpha1.Cronjob, opts v1.CreateOptions) (result *v1alpha1.Cronjob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cronjobsResource, c.ns, cronjob), &v1alpha1.Cronjob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cronjob), err
}

// Update takes the representation of a cronjob and updates it. Returns the server's representation of the cronjob, and an error, if there is any.
func (c *FakeCronjobs) Update(ctx context.Context, cronjob *v1alpha1.Cronjob, opts v1.UpdateOptions) (result *v1alpha1.Cronjob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cronjobsResource, c.ns, cronjob), &v1alpha1.Cronjob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cronjob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCronjobs) UpdateStatus(ctx context.Context, cronjob *v1alpha1.Cronjob, opts v1.UpdateOptions) (*v1alpha1.Cronjob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cronjobsResource, "status", c.ns, cronjob), &v1alpha1.Cronjob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cronjob), err
}

// Delete takes name of the cronjob and deletes it. Returns an error if one occurs.
func (c *FakeCronjobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cronjobsResource, c.ns, name), &v1alpha1.Cronjob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCronjobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cronjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CronjobList{})
	return err
}

// Patch applies the patch and returns the patched cronjob.
func (c *FakeCronjobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Cronjob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cronjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Cronjob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cronjob), err
}
