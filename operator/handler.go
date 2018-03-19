package operator

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	"github.com/marjoram/duke-crd/log"
)

// Handler is the pipeline event handler
type handler struct {
	dukeService duked.Syncer
	logger      log.Logger
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
	pipeline, ok := obj.(*dukedv1alpha1.Pipeline)
	if !ok {
		return fmt.Errorf("%v is not a duke object", obj.GetObjectKind())
	}

	return h.dukeService.EnsureDukePipeline(duke)
}

// Delete will ensure the reunited pod yadyada
func (h *handler) Delete(name string) error {
	return h.dukeService.DeletePod(name)
}
