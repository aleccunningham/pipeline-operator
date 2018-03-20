package operator

import (
	"github.com/spotahome/kooper/client/crd"
	"github.com/spotahome/kooper/operator"
	"github.com/spotahome/kooper/operator/controller"
	"k8s.io/client-go/kubernetes"

	agentdukev1alpha1 "github.com/marjoram/pipeline-operator/apis/agent.cncd.io/v1alpha1"
	"github.com/marjoram/pipeline-operator/log"
)

func New(cfg Config, ageCli agentdukev1alpha1.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface, logger log.logger) (operator.Operator, error) {
	// Create the CRD
	AgentCRD := newAgennt(ageCli, crdCli, kubeCli)
	// Create the handler
	handler := newHandler(kubeCli, logger)
	// Create controller
	ctrl := controller.NewSequential(cfg.ResyncPeriod, handler, pipelineCRD, nil, logger)
	// Assemble CRD and controller to create the operator
	return operator.NewOperator(ageCRD, ctrl, logger), nil
}
