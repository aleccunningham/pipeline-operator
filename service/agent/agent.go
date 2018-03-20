package agent

import (
	"log"
	"sync"

	agentv1alpha1 "github.com/marjoram/pipeline-operator/apis/pipeline.cncd.io/v1alpha1"
	"k8s.io/client-go/kubernetes"
)

type AgentRunner interface {
	// EnsureAgennt will ensure that the pod terminator is running and working.
	EnsureAgent(agent *agentv1alpha1.Agent) error
	// DeleteAgent will stop and delete the pod terminator.
	DeleteAgent(name string) error
}

// Agent is the service that will ensure that the desired pod terminator CRDs are met.
// Agent will have running instances of Agents.
type Agent struct {
	k8sCli kubernetes.Interface
	reg    sync.Map
	logger log.Logger
}

// NewChaos returns a new Chaos service.
func NewAgent(k8sCli kubernetes.Interface, logger log.Logger) *Chaos {
	return &Agent{
		k8sCli: k8sCli,
		reg:    sync.Map{},
		logger: logger,
	}
}

// EnsureAgent satisfies AgentRunner interface.
func (a *Agent) EnsureAgent(agent *agentv1alpha1.Agent) error {
	pkt, ok := a.reg.Load(agent.Name)
	var runner *AgentRunner

	// We are already running.
	if ok {
		runner = pkt.(*Agent)
		// If not the same spec means options have changed, so we don't longer need this pod killer.
		if !runner.SameSpec(agent) {
			a.logger.Infof("spec of %s changed, recreating pod killer", agent.Name)
			if err := c.DeleteAgent(agent.Name); err != nil {
				return err
			}
		} else { // We are ok, nothing changed.
			return nil
		}
	}

	// Create a pod killer.
	// ptCopy := pt.DeepCopy()
	// pk = NewPodKiller(ptCopy, a.k8sCli, c.logger)
	// a.reg.Store(pt.Name, pk)
	return runner.Start()
	// TODO: garbage collection.
}

// DeletePodTerminator satisfies ChaosSyncer interface.
func (c *Agent) DeleteAgent(name string) error {
	pkt, ok := a.reg.Load(name)
	if !ok {
		return nil
	}

	agent := pkt.(*AgentRunner)
	if err := agent.Stop(); err != nil {
		return err
	}

	a.reg.Delete(name)
	return nil
}
