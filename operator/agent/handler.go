package agent

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	agentv1alpha1 "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned/typed/agent.cncd.io/v1alpha1"
	"github.com/marjoram/pipeline-operator/log"
	service "github.com/marjoram/pipeline-operator/service/agent"
)

// Handler is the agentn event handler
type handler struct {
	agentService service.AgentClient
	logger       log.Logger
}

// newHandler returns a new handler
func newHandler(k8sCli kubernetes.Interface, logger log.Logger) *handler {
	return &handler{
		pipelineService: agentv1alpha1.NewAgent(k8sCli, logger),
		logger:          logger,
	}
}

// Add will ensure that the required pipeline workers are running
func (h *handler) Add(obj runtime.Object) error {
	agent, ok := obj.(*agentv1alpha1.Agent)
	if !ok {
		return fmt.Errorf("%v is not a agent object", obj.GetObjectKind())
	}

	return h.agentService.Healthcheck()
}

// Delete will ensure the reunited pod yadyada
func (h *handler) Delete(name string) error {
	return h.agentService.Healthcheck()
}
