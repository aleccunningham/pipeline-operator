package fake

import (
	v1alpha1 "github.com/marjoram/duke-crd/client/k8s/clientset/versioned/typed/duke/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDukeV1alpha1 struct {
	*testing.Fake
}

func (c *FakeDukeV1alpha1) Pipelines() v1alpha1.PipelineInterface {
	return &FakePipelines{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDukeV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
