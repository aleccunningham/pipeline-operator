package agent

import (
	"github.com/spotahome/kooper/client/crd"
	"github.com/spotahome/kooper/operator"
	"github.com/spotahome/kooper/operator/controller"
	"k8s.io/client-go/kubernetes"

	agentk8sCli "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned"
	"github.com/marjoram/pipeline-operator/log"
	o "github.com/marjoram/pipeline-operator/operator/agent"
)

func New(cfg Config, ageCli agentk8sCli.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface, logger log.logger) (o.Operator, error) {
	// Create the CRD
	agentCRD := newAgentCRD(ageCli, crdCli, kubeCli)
	// Create the handler
	handler := newHandler(kubeCli, logger)
	// Create controller
	ctrl := controller.NewSequential(cfg.ResyncPeriod, handler, agentCRD, nil, logger)
	// Assemble CRD and controller to create the operator
	return operator.NewOperator(ageCRD, ctrl, logger), nil
}
