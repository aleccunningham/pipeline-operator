package operator

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	agentdukev1alpha1 "github.com/marjoram/pipeline-operator/apis/agent.duke.lol/v1alpha1"
	"github.com/marjoram/pipeline-operator/log"
)

// Handler is the pipeline event handler
type handler struct {
	agentService agentdukev1alpha1.Syncer
	logger       log.Logger
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
	agent, ok := obj.(*agentdukev1alpha1.Agennt)
	if !ok {
		return fmt.Errorf("%v is not a agent object", obj.GetObjectKind())
	}

	return h.agentService.EnsureAgent(agent)
}

// Delete will ensure the reunited pod yadyada
func (h *handler) Delete(name string) error {
	return h.agentService.DeletePod(name)
}