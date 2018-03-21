package pipeline

import (
	"github.com/spotahome/kooper/client/crd"
	"github.com/spotahome/kooper/operator"
	"github.com/spotahome/kooper/operator/controller"
	"k8s.io/client-go/kubernetes"

	pipelinek8sCli "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned"
	"github.com/marjoram/pipeline-operator/log"
	o "github.com/marjoram/pipeline-operator/operator"
)

func New(cfg Config, pipeCli pipelinek8sCli.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface, logger log.Logger) (o.pipeline.Operator, error) {
	// Create the CRD
	pipelineCRD := newPipelineCRD(pipeCli, crdCli, kubeCli)
	// Create the handler
	handler := newHandler(kubeCli, logger)
	// Create controller
	ctrl := controller.NewSequential(cfg.ResyncPeriod, handler, pipelineCRD, nil, logger)
	// Assemble CRD and controller to create the operator
	return operator.NewOperator(pipelineCRD, ctrl, logger), nil
}
