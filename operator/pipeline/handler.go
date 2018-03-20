package operator

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	pipelinedukev1alpha1 "github.com/marjoram/pipeline-operator/apis/pipeline.duke.lol/v1alpha1"
	"github.com/marjoram/pipeline-operator/log"
)

// Handler is the pipeline event handler
type handler struct {
	pipelineService pipelinedukev1alpha1.Syncer
	logger          log.Logger
}

// newHandler returns a new handler
func newHandler(k8sCli kubernetes.Interface, logger log.Logger) *handler {
	return &handler{
		pipelineService: pipelinedukev1alpha1.NewPipeline(k8sCli, logger),
		logger:          logger,
	}
}

// Add will ensure that the required pipeline workers are runninng
func (h *handler) Add(obj runtime.Object) error {
	pipeline, ok := obj.(*pipelinedukev1alpha1.Pipeline)
	if !ok {
		return fmt.Errorf("%v is not a duke object", obj.GetObjectKind())
	}

	return h.pipelineService.EnsurePipeline(pipeline)
}

// Delete will ensure the reunited pod yadyada
func (h *handler) Delete(name string) error {
	return h.pipelineService.DeletePod(name)
}
