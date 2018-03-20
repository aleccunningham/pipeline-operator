package v1alpha1

import (
	v1alpha1 "github.com/marjoram/pipeline-operator/apis/autoscaler.duke.lol/v1alpha1"
	scheme "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AutoScalersGetter has a method to return a AutoScalerInterface.
// A group's client should implement this interface.
type AutoScalersGetter interface {
	AutoScalers() AutoScalerInterface
}

// AutoScalerInterface has methods to work with AutoScaler resources.
type AutoScalerInterface interface {
	Create(*v1alpha1.AutoScaler) (*v1alpha1.AutoScaler, error)
	Update(*v1alpha1.AutoScaler) (*v1alpha1.AutoScaler, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AutoScaler, error)
	List(opts v1.ListOptions) (*v1alpha1.AutoScalerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoScaler, err error)
	AutoScalerExpansion
}

// autoScalers implements AutoScalerInterface
type autoScalers struct {
	client rest.Interface
}

// newAutoScalers returns a AutoScalers
func newAutoScalers(c *AgentV1alpha1Client) *autoScalers {
	return &autoScalers{
		client: c.RESTClient(),
	}
}

// Get takes name of the autoScaler, and returns the corresponding autoScaler object, and an error if there is any.
func (c *autoScalers) Get(name string, options v1.GetOptions) (result *v1alpha1.AutoScaler, err error) {
	result = &v1alpha1.AutoScaler{}
	err = c.client.Get().
		Resource("autoscalers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AutoScalers that match those selectors.
func (c *autoScalers) List(opts v1.ListOptions) (result *v1alpha1.AutoScalerList, err error) {
	result = &v1alpha1.AutoScalerList{}
	err = c.client.Get().
		Resource("autoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested autoScalers.
func (c *autoScalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("autoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a autoScaler and creates it.  Returns the server's representation of the autoScaler, and an error, if there is any.
func (c *autoScalers) Create(autoScaler *v1alpha1.AutoScaler) (result *v1alpha1.AutoScaler, err error) {
	result = &v1alpha1.AutoScaler{}
	err = c.client.Post().
		Resource("autoscalers").
		Body(autoScaler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a autoScaler and updates it. Returns the server's representation of the autoScaler, and an error, if there is any.
func (c *autoScalers) Update(autoScaler *v1alpha1.AutoScaler) (result *v1alpha1.AutoScaler, err error) {
	result = &v1alpha1.AutoScaler{}
	err = c.client.Put().
		Resource("autoscalers").
		Name(autoScaler.Name).
		Body(autoScaler).
		Do().
		Into(result)
	return
}

// Delete takes name of the autoScaler and deletes it. Returns an error if one occurs.
func (c *autoScalers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("autoscalers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *autoScalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("autoscalers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched autoScaler.
func (c *autoScalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoScaler, err error) {
	result = &v1alpha1.AutoScaler{}
	err = c.client.Patch(pt).
		Resource("autoscalers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
