package agent

import (
	"log"
	"sync"

	"k8s.io/client-go/kubernetes"
)

type AgentClient interface {
	Healthcheck() error
}

// Agent is the service that will ensure that the desired pod terminator CRDs are met.
// Agent will have running instances of Agents.
type Agent struct {
	k8sCli kubernetes.Interface
	reg    sync.Map
	logger log.Logger
}

// NewChaos returns a new Chaos service.
func NewAgent(k8sCli kubernetes.Interface, logger log.Logger) *Agent {
	return &Agent{
		k8sCli: k8sCli,
		reg:    sync.Map{},
		logger: logger,
	}
}

func (a *Agent) Healthcheck() error {
	return nil
}
