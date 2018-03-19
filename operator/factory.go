package operator

import (
	"github.com/spotahome/kooper/client/crd"
	"github.com/spotahome/kooper/operator"
	"github.com/spotahome/kooper/operator/controller"
	"k8s.io/client-go/kubernetes"

	podtermk8scli "github.com/spotahome/kooper/examples/pod-terminator-operator/client/k8s/clientset/versioned"
	"github.com/spotahome/kooper/examples/pod-terminator-operator/log"
)

func New(cfg Config, dukedCli.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface, logger log.logger) (operator.Operator, error) {
		// Create the CRD
		dukedCRD := newPipeline(dukedCli, crdCli, kubeCli)
		// Create the handler
		handler := newHandler(kubeCli, logger)
		// Create controller
		ctrl := controller.NewSequential(cfg.ResyncPeriod, handler, dukedCRD, nil, logger)
		// Assemble CRD and controller to create the operator
		return operator.NewOperator(dukedCRD, ctrl, logger), nil
}
