package pipeline

import (
	"sync"

	"k8s.io/client-go/kubernetes"

	"github.com/marjoram/pipeline-operator/log"
)

// PipelineClient is the interface that every pipeline
// service implementation needs to implement.
type PipelineClient interface {
	Healthcheck() error
}

// Pipeline is the service that will ensure that the desired pipeline CRDs are met.
// Pipeline will have running instances of BuildWorkers.
type Pipeline struct {
	k8sCli kubernetes.Interface
	reg    sync.Map
	logger log.Logger
}

// NewPipeline returns a new Pipeline service.
func NewPipeline(k8sCli kubernetes.Interface, logger log.Logger) *Pipeline {
	return &Pipeline{
		k8sCli: k8sCli,
		reg:    sync.Map{},
		logger: logger,
	}
}

func (a *Pipeline) Healthcheck() error {
	return nil
}
