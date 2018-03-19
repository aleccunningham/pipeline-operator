package operator

import (
	"github.com/spotahome/kooper/client/crd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	dukedv1alpha1 "github.com/marjoram/duked/apis/duked/v1alpha1"
)

// dukedPipeline is a Pipeline CRD
type dukedPipeline struct {
	crdCli   crd.Interface
	kubecCli kubernetes.Interface
	dukeCli  pipeline.Interface
}

func newDukedPipelineCRD(dukeCli Pipeline.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface) *podTerminatorCRD {
	return &podTerminatorCRD{
		crdCli:   crdCli,
		dukeCli:  dukeCli,
		kubecCli: kubeCli,
	}
}

// podTerminatorCRD satisfies resource.crd interface.
func (p *dukedPipeline) Initialize() error {
	crd := crd.Conf{
		Kind:       dukedv1alpha1.PipelineKind,
		NamePlural: dukedv1alpha1.PipelineName,
		Group:      dukedv1alpha1.SchemeGroupVersion.Group,
		Version:    dukedv1alpha1.SchemeGroupVersion.Version,
		Scope:      dukedv1alpha1.PipelineScope,
	}

	return p.crdCli.EnsurePresent(crd)
}

// GetListerWatcher satisfies resource.crd interface (and retrieve.Retriever).
func (p *dukedPipeline) GetListerWatcher() cache.ListerWatcher {
	return &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return p.podTermCli.DV1alpha1().PodTerminators().List(options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return p.podTermCli.ChaosV1alpha1().PodTerminators().Watch(options)
		},
	}
}

// GetObject satisfies resource.crd interface (and retrieve.Retriever).
func (p *podTerminatorCRD) GetObject() runtime.Object {
	return &chaosv1alpha1.PodTerminator{}
}
