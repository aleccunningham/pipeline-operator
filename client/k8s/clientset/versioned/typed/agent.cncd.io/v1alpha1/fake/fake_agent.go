package fake

import (
	v1alpha1 "github.com/marjoram/pipeline-operator/apis/agent.cncd.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAgents implements AgentInterface
type FakeAgents struct {
	Fake *FakeAgentV1alpha1
}

var agentsResource = schema.GroupVersionResource{Group: "agent.cncd.io", Version: "v1alpha1", Resource: "agents"}

var agentsKind = schema.GroupVersionKind{Group: "agent.cncd.io", Version: "v1alpha1", Kind: "Agent"}

// Get takes name of the agent, and returns the corresponding agent object, and an error if there is any.
func (c *FakeAgents) Get(name string, options v1.GetOptions) (result *v1alpha1.Agent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(agentsResource, name), &v1alpha1.Agent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Agent), err
}

// List takes label and field selectors, and returns the list of Agents that match those selectors.
func (c *FakeAgents) List(opts v1.ListOptions) (result *v1alpha1.AgentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(agentsResource, agentsKind, opts), &v1alpha1.AgentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AgentList{}
	for _, item := range obj.(*v1alpha1.AgentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested agents.
func (c *FakeAgents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(agentsResource, opts))
}

// Create takes the representation of a agent and creates it.  Returns the server's representation of the agent, and an error, if there is any.
func (c *FakeAgents) Create(agent *v1alpha1.Agent) (result *v1alpha1.Agent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(agentsResource, agent), &v1alpha1.Agent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Agent), err
}

// Update takes the representation of a agent and updates it. Returns the server's representation of the agent, and an error, if there is any.
func (c *FakeAgents) Update(agent *v1alpha1.Agent) (result *v1alpha1.Agent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(agentsResource, agent), &v1alpha1.Agent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Agent), err
}

// Delete takes name of the agent and deletes it. Returns an error if one occurs.
func (c *FakeAgents) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(agentsResource, name), &v1alpha1.Agent{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAgents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(agentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AgentList{})
	return err
}

// Patch applies the patch and returns the patched agent.
func (c *FakeAgents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Agent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(agentsResource, name, data, subresources...), &v1alpha1.Agent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Agent), err
}
