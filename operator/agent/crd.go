package operator

import (
	"github.com/spotahome/kooper/client/crd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	agentdukev1alpha1 "github.com/marjoram/pipeline-operator/apis/agent.duke.lol/v1alpha1"
)

// AgentCRD is a Pipeline CRD
type AgentCRD struct {
	crdCli   crd.Interface
	kubecCli kubernetes.Interface
	ageCli   agentdukev1alpha1.Interface
}

func newPipelineCRD(ageCli agentdukev1alpha1.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface) *PipelineCRD {
	return &PipelineCRD{
		crdCli:   crdCli,
		ageCli:   ageCli,
		kubecCli: kubeCli,
	}
}

// podTerminatorCRD satisfies resource.crd interface.
func (a *AgentCRD) Initialize() error {
	crd := crd.Conf{
		Kind:       pipelinedukev1alpha1.AgentKind,
		NamePlural: pipelinedukev1alpha1.AgentName,
		Group:      pipelinedukev1alpha1.SchemeGroupVersion.Group,
		Version:    pipelinedukev1alpha1.SchemeGroupVersion.Version,
		Scope:      pipelinedukev1alpha1.AgentScope,
	}

	return a.crdCli.EnsurePresent(crd)
}

// GetListerWatcher satisfies resource.crd interface (and retrieve.Retriever).
func (a *AgentCRD) GetListerWatcher() cache.ListerWatcher {
	return &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return a.ageCli.ListAgents("", options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return a.ageCli.WatchAgents("", options)
		},
	}
}

// GetObject satisfies resource.crd interface (and retrieve.Retriever).
func (a *AgentCRD) GetObject() runtime.Object {
	return &pipelinedukev1alpha1.Agent{}
}
