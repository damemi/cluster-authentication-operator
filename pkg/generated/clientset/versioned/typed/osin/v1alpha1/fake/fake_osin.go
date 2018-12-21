// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/openshift/cluster-osin-operator/pkg/apis/osin/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOsins implements OsinInterface
type FakeOsins struct {
	Fake *FakeOsinV1alpha1
	ns   string
}

var osinsResource = schema.GroupVersionResource{Group: "osin.openshift.io", Version: "v1alpha1", Resource: "osins"}

var osinsKind = schema.GroupVersionKind{Group: "osin.openshift.io", Version: "v1alpha1", Kind: "Osin"}

// Get takes name of the osin, and returns the corresponding osin object, and an error if there is any.
func (c *FakeOsins) Get(name string, options v1.GetOptions) (result *v1alpha1.Osin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(osinsResource, c.ns, name), &v1alpha1.Osin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Osin), err
}

// List takes label and field selectors, and returns the list of Osins that match those selectors.
func (c *FakeOsins) List(opts v1.ListOptions) (result *v1alpha1.OsinList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(osinsResource, osinsKind, c.ns, opts), &v1alpha1.OsinList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OsinList{ListMeta: obj.(*v1alpha1.OsinList).ListMeta}
	for _, item := range obj.(*v1alpha1.OsinList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested osins.
func (c *FakeOsins) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(osinsResource, c.ns, opts))

}

// Create takes the representation of a osin and creates it.  Returns the server's representation of the osin, and an error, if there is any.
func (c *FakeOsins) Create(osin *v1alpha1.Osin) (result *v1alpha1.Osin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(osinsResource, c.ns, osin), &v1alpha1.Osin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Osin), err
}

// Update takes the representation of a osin and updates it. Returns the server's representation of the osin, and an error, if there is any.
func (c *FakeOsins) Update(osin *v1alpha1.Osin) (result *v1alpha1.Osin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(osinsResource, c.ns, osin), &v1alpha1.Osin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Osin), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOsins) UpdateStatus(osin *v1alpha1.Osin) (*v1alpha1.Osin, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(osinsResource, "status", c.ns, osin), &v1alpha1.Osin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Osin), err
}

// Delete takes name of the osin and deletes it. Returns an error if one occurs.
func (c *FakeOsins) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(osinsResource, c.ns, name), &v1alpha1.Osin{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOsins) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(osinsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OsinList{})
	return err
}

// Patch applies the patch and returns the patched osin.
func (c *FakeOsins) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Osin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(osinsResource, c.ns, name, data, subresources...), &v1alpha1.Osin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Osin), err
}