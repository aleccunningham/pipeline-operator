package chaos

import (
	"sync"

	"k8s.io/client-go/kubernetes"

	pipelinedukev1alpha1 "github.com/marjoram/pipeline-operator/apis/pipeline.duke.lol/v1alpha1"
	"github.com/marjoram/pipeline-operator/log"
)

// Syncer is the interface that every chaos service implementation
// needs to implement.
type Syncer interface {
	// EnsurePipeline will ensure that the pipeline is running and working.
	EnsurePipeline(pt *pipelinedukev1alpha1.Pipeline) error
	// DeletePipeline will stop and delete the pipeline
	DeletePipeline(name string) error
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

// EnsurePipeline satisfies Syncer interface.
func (c *Chaos) EnsurePipeline(pt *pipelinedukev1alpha1.Pipeline) error {
	pkt, ok := c.reg.Load(pt.Name)
	var pk *PodKiller

	// We are already running.
	if ok {
		pk = pkt.(*PodKiller)
		// If not the same spec means options have changed, so we don't longer need this pod killer.
		if !pk.SameSpec(pt) {
			c.logger.Infof("spec of %s changed, recreating pod killer", pt.Name)
			if err := c.DeletePodTerminator(pt.Name); err != nil {
				return err
			}
		} else { // We are ok, nothing changed.
			return nil
		}
	}

	// Create a pod killer.
	ptCopy := pt.DeepCopy()
	pk = NewPodKiller(ptCopy, c.k8sCli, c.logger)
	c.reg.Store(pt.Name, pk)
	return pk.Start()
	// TODO: garbage collection.
}
