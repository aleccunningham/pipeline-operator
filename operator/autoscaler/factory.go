package operator

import (
	"github.com/spotahome/kooper/client/crd"
	"github.com/spotahome/kooper/operator"
	"github.com/spotahome/kooper/operator/controller"
	"k8s.io/client-go/kubernetes"

	autoscalerdukev1alpha1 "github.com/marjoram/pipeline-operator/apis/autoscaler.duke.lol/v1alpha1"
	"github.com/marjoram/pipeline-operator/log"
)

func New(cfg Config, scaleCli autoscalerdukev1alpha1.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface, logger log.logger) (operator.Operator, error) {
	// Create the CRD
	autoScalerCRD := newPipeline(scaleCli, crdCli, kubeCli)
	// Create the handler
	handler := newHandler(kubeCli, logger)
	// Create controller
	ctrl := controller.NewSequential(cfg.ResyncPeriod, handler, scaleCRD, nil, logger)
	// Assemble CRD and controller to create the operator
	return operator.NewOperator(autoScalerCRD, ctrl, logger), nil
}
