package fake

import (
	v1alpha1 "github.com/marjoram/pipeline-operator/apis/autoscaler.duke.lol/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAutoScalers implements AutoScalerInterface
type FakeAutoScalers struct {
	Fake *FakeAutoscalerV1alpha1
}

var autoscalersResource = schema.GroupVersionResource{Group: "autoscaler.duke.lol", Version: "v1alpha1", Resource: "autoscalers"}

var autoscalersKind = schema.GroupVersionKind{Group: "autoscaler.duke.lol", Version: "v1alpha1", Kind: "AutoScaler"}

// Get takes name of the autoScaler, and returns the corresponding autoScaler object, and an error if there is any.
func (c *FakeAutoScalers) Get(name string, options v1.GetOptions) (result *v1alpha1.AutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(autoscalersResource, name), &v1alpha1.AutoScaler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoScaler), err
}

// List takes label and field selectors, and returns the list of AutoScalers that match those selectors.
func (c *FakeAutoScalers) List(opts v1.ListOptions) (result *v1alpha1.AutoScalerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(autoscalersResource, autoscalersKind, opts), &v1alpha1.AutoScalerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutoScalerList{}
	for _, item := range obj.(*v1alpha1.AutoScalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested autoScalers.
func (c *FakeAutoScalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(autoscalersResource, opts))
}

// Create takes the representation of a autoScaler and creates it.  Returns the server's representation of the autoScaler, and an error, if there is any.
func (c *FakeAutoScalers) Create(autoScaler *v1alpha1.AutoScaler) (result *v1alpha1.AutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(autoscalersResource, autoScaler), &v1alpha1.AutoScaler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoScaler), err
}

// Update takes the representation of a autoScaler and updates it. Returns the server's representation of the autoScaler, and an error, if there is any.
func (c *FakeAutoScalers) Update(autoScaler *v1alpha1.AutoScaler) (result *v1alpha1.AutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(autoscalersResource, autoScaler), &v1alpha1.AutoScaler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoScaler), err
}

// Delete takes name of the autoScaler and deletes it. Returns an error if one occurs.
func (c *FakeAutoScalers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(autoscalersResource, name), &v1alpha1.AutoScaler{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutoScalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(autoscalersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutoScalerList{})
	return err
}

// Patch applies the patch and returns the patched autoScaler.
func (c *FakeAutoScalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(autoscalersResource, name, data, subresources...), &v1alpha1.AutoScaler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoScaler), err
}
