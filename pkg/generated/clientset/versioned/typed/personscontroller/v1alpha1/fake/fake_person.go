/*
Copyright The Kubernetes Authors.

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

	v1alpha1 "github.com/andreaskaris/people-controller/pkg/apis/personscontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePersons implements PersonInterface
type FakePersons struct {
	Fake *FakePeopleV1alpha1
	ns   string
}

var personsResource = schema.GroupVersionResource{Group: "people.example.com", Version: "v1alpha1", Resource: "persons"}

var personsKind = schema.GroupVersionKind{Group: "people.example.com", Version: "v1alpha1", Kind: "Person"}

// Get takes name of the person, and returns the corresponding person object, and an error if there is any.
func (c *FakePersons) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Person, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(personsResource, c.ns, name), &v1alpha1.Person{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Person), err
}

// List takes label and field selectors, and returns the list of Persons that match those selectors.
func (c *FakePersons) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PersonList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(personsResource, personsKind, c.ns, opts), &v1alpha1.PersonList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PersonList{ListMeta: obj.(*v1alpha1.PersonList).ListMeta}
	for _, item := range obj.(*v1alpha1.PersonList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested persons.
func (c *FakePersons) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(personsResource, c.ns, opts))

}

// Create takes the representation of a person and creates it.  Returns the server's representation of the person, and an error, if there is any.
func (c *FakePersons) Create(ctx context.Context, person *v1alpha1.Person, opts v1.CreateOptions) (result *v1alpha1.Person, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(personsResource, c.ns, person), &v1alpha1.Person{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Person), err
}

// Update takes the representation of a person and updates it. Returns the server's representation of the person, and an error, if there is any.
func (c *FakePersons) Update(ctx context.Context, person *v1alpha1.Person, opts v1.UpdateOptions) (result *v1alpha1.Person, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(personsResource, c.ns, person), &v1alpha1.Person{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Person), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePersons) UpdateStatus(ctx context.Context, person *v1alpha1.Person, opts v1.UpdateOptions) (*v1alpha1.Person, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(personsResource, "status", c.ns, person), &v1alpha1.Person{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Person), err
}

// Delete takes name of the person and deletes it. Returns an error if one occurs.
func (c *FakePersons) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(personsResource, c.ns, name, opts), &v1alpha1.Person{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePersons) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(personsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PersonList{})
	return err
}

// Patch applies the patch and returns the patched person.
func (c *FakePersons) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Person, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(personsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Person{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Person), err
}
