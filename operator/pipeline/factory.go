package operator

import (
	"github.com/spotahome/kooper/client/crd"
	"github.com/spotahome/kooper/operator"
	"github.com/spotahome/kooper/operator/controller"
	"k8s.io/client-go/kubernetes"

	pipelinedukev1alpha1 "github.com/marjoram/pipeline-operator/apis/pipeline.duke.lol/v1alpha1"
	"github.com/marjoram/pipeline-operator/log"
)

func New(cfg Config, pipeCli pipelinedukev1alpha1.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface, logger log.logger) (operator.Operator, error) {
	// Create the CRD
	pipelineCRD := newPipeline(pipeCli, crdCli, kubeCli)
	// Create the handler
	handler := newHandler(kubeCli, logger)
	// Create controller
	ctrl := controller.NewSequential(cfg.ResyncPeriod, handler, pipelineCRD, nil, logger)
	// Assemble CRD and controller to create the operator
	return operator.NewOperator(pipelineCRD, ctrl, logger), nil
}
