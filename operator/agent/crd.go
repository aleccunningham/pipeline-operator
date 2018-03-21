package agent

import (
	"github.com/spotahome/kooper/client/crd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	agentv1alpha1 "github.com/marjoram/pipeline-operator/apis/agent.cncd.io/v1alpha1"
	agentk8sCli "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned"
)

// AgentCRD is a Pipeline CRD
type AgentCRD struct {
	crdCli   crd.Interface
	kubecCli kubernetes.Interface
	ageCli   agentk8sCli.Interface
}

func newAgentCRD(ageCli agentk8sCli.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface) *AgentCRD {
	return &AgentCRD{
		crdCli:   crdCli,
		ageCli:   ageCli,
		kubecCli: kubeCli,
	}
}

// agentCRD satisfies resource.crd interface.
func (a *AgentCRD) Initialize() error {
	crd := crd.Conf{
		Kind:       agentv1alpha1.AgentKind,
		NamePlural: agentv1alpha1.AgentName,
		Group:      agentv1alpha1.SchemeGroupVersion.Group,
		Version:    agentv1alpha1.SchemeGroupVersion.Version,
		Scope:      agentv1alpha1.AgentScope,
	}

	return a.crdCli.EnsurePresent(crd)
}

// GetListerWatcher satisfies resource.crd interface (and retrieve.Retriever).
func (a *AgentCRD) GetListerWatcher() cache.ListerWatcher {
	return &cache.ListWatch{
		ListFunc: func(opts metav1.ListOptions) (runtime.Object, error) {
			return a.ageCli.AgentV1alpha1().Agents().List(opts)
		},
		WatchFunc: func(opts metav1.ListOptions) (watch.Interface, error) {
			return a.ageCli.AgentV1alpha1().Agents().Watch(opts)
		},
	}
}

// GetObject satisfies resource.crd interface (and retrieve.Retriever).
func (a *AgentCRD) GetObject() runtime.Object {
	return &agentv1alpha1.Agent{}
}
