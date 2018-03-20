package fake

import (
	clientset "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned"
	agentv1alpha1 "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned/typed/agent.duke.lol/v1alpha1"
	fakeagentv1alpha1 "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned/typed/agent.duke.lol/v1alpha1/fake"
	autoscalerv1alpha1 "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned/typed/autoscaler.duke.lol/v1alpha1"
	fakeautoscalerv1alpha1 "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned/typed/autoscaler.duke.lol/v1alpha1/fake"
	pipelinev1alpha1 "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned/typed/pipeline.duke.lol/v1alpha1"
	fakepipelinev1alpha1 "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned/typed/pipeline.duke.lol/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", testing.DefaultWatchReactor(watch.NewFake(), nil))

	return &Clientset{fakePtr, &fakediscovery.FakeDiscovery{Fake: &fakePtr}}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// AgentV1alpha1 retrieves the AgentV1alpha1Client
func (c *Clientset) AgentV1alpha1() agentv1alpha1.AgentV1alpha1Interface {
	return &fakeagentv1alpha1.FakeAgentV1alpha1{Fake: &c.Fake}
}

// Agent retrieves the AgentV1alpha1Client
func (c *Clientset) Agent() agentv1alpha1.AgentV1alpha1Interface {
	return &fakeagentv1alpha1.FakeAgentV1alpha1{Fake: &c.Fake}
}

// AutoscalerV1alpha1 retrieves the AutoscalerV1alpha1Client
func (c *Clientset) AutoscalerV1alpha1() autoscalerv1alpha1.AutoscalerV1alpha1Interface {
	return &fakeautoscalerv1alpha1.FakeAutoscalerV1alpha1{Fake: &c.Fake}
}

// Autoscaler retrieves the AutoscalerV1alpha1Client
func (c *Clientset) Autoscaler() autoscalerv1alpha1.AutoscalerV1alpha1Interface {
	return &fakeautoscalerv1alpha1.FakeAutoscalerV1alpha1{Fake: &c.Fake}
}

// PipelineV1alpha1 retrieves the PipelineV1alpha1Client
func (c *Clientset) PipelineV1alpha1() pipelinev1alpha1.PipelineV1alpha1Interface {
	return &fakepipelinev1alpha1.FakePipelineV1alpha1{Fake: &c.Fake}
}

// Pipeline retrieves the PipelineV1alpha1Client
func (c *Clientset) Pipeline() pipelinev1alpha1.PipelineV1alpha1Interface {
	return &fakepipelinev1alpha1.FakePipelineV1alpha1{Fake: &c.Fake}
}
