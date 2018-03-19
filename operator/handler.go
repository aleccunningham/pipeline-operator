package operator

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	"github.com/spotahome/kooper/examples/pod-terminator-operator/log"
)

// Handler is the pipeline event handler
type handler struct {
	dukedService duked.Syncer
	logger       log.Logger
}

// newHandler returns a new handler
func newHandler(k8sCli kubernetes.Interface, logger log.Logger) *handler {
	return &handler{
		dukedService: duked.NewPipeline(k8sCli, logger),
		logger:       logger,
	}
}

// Add will ensure that the required pipeline workers are runninng
func (h *handler) Add(obj runtime.Object) error {
	duked, ok := obj.(*dukedv1alpha1.Pipeline)
	if !ok {
		return fmt.Errorf("%v is not a duked object", obj.GetObjectKind())
	}

	return h.dukedService.EnsureDukedPipeline(duked)
}

// Delete will ensure the reunited pod yadyada
func (h *handler) Delete(name string) error {
	return h.dukedService.DeletePod(name)
}
