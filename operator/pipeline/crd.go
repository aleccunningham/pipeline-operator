package operator

import (
	"github.com/spotahome/kooper/client/crd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	pipelinev1alpha1 "github.com/marjoram/pipeline-operator/apis/pipeline.cncd.io/v1alpha1"
)

// PipelineCRD is a Pipeline CRD
type PipelineCRD struct {
	crdCli   crd.Interface
	kubecCli kubernetes.Interface
	pipeCli  pipeline.Interface
}

func newPipelineCRD(pipeCli pipelinev1alpha1.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface) *PipelineCRD {
	return &PipelineCRD{
		crdCli:   crdCli,
		pipeCli:  pipeCli,
		kubecCli: kubeCli,
	}
}

// podTerminatorCRD satisfies resource.crd interface.
func (p *PipelineCRD) Initialize() error {
	crd := crd.Conf{
		Kind:       pipelinev1alpha1.PipelineKind,
		NamePlural: pipelinev1alpha1.PipelineName,
		Group:      pipelinev1alpha1.SchemeGroupVersion.Group,
		Version:    pipelinev1alpha1.SchemeGroupVersion.Version,
		Scope:      pipelinev1alpha1.PipelineScope,
	}

	return p.crdCli.EnsurePresent(crd)
}

// GetListerWatcher satisfies resource.crd interface (and retrieve.Retriever).
func (p *PipelineCRD) GetListerWatcher() cache.ListerWatcher {
	return &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return p.pipeCli.ListPipelines("", options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return p.pipeCli.WatchPipelines("", options)
		},
	}
}

// GetObject satisfies resource.crd interface (and retrieve.Retriever).
func (p *podTerminatorCRD) GetObject() runtime.Object {
	return &pipelinev1alpha1.Pipeline{}
}
