package fake

import (
	v1alpha1 "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned/typed/autoscaler.duke.lol/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAutoscalerV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAutoscalerV1alpha1) AutoScalers() v1alpha1.AutoScalerInterface {
	return &FakeAutoScalers{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAutoscalerV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
