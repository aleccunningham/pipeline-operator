package operator

import (
	"github.com/spotahome/kooper/client/crd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	autoscalerdukev1alpha1 "github.com/marjoram/pipeline-operator/apis/autoscaler.duke.lol/v1alpha1"
)

// AutoScalerCRD is a Pipeline CRD
type AutoScalerCRD struct {
	crdCli   crd.Interface
	kubecCli kubernetes.Interface
	pipeCli  pipeline.Interface
}

func newAutoScalerCRD(scaleCli AutoScaler.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface) *PipelineCRD {
	return &PipelineCRD{
		crdCli:   crdCli,
		scaleCli: scaleCli,
		kubecCli: kubeCli,
	}
}

// podTerminatorCRD satisfies resource.crd interface.
func (a *AutoScalerCRD) Initialize() error {
	crd := crd.Conf{
		Kind:       autoscalerdukev1alpha1.AutoScalerKind,
		NamePlural: autoscalerdukev1alpha1.AutoScalerName,
		Group:      autoscalerdukev1alpha1.SchemeGroupVersion.Group,
		Version:    autoscalerdukev1alpha1.SchemeGroupVersion.Version,
		Scope:      autoscalerdukev1alpha1.AutoScalerScope,
	}

	return a.crdCli.EnsurePresent(crd)
}

// GetListerWatcher satisfies resource.crd interface (and retrieve.Retriever).
func (a *AutoScalerCRD) GetListerWatcher() cache.ListerWatcher {
	return &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return a.scaleCli.ListAutoScalers("", options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return a.scaleCli.WatchAutoScalers("", options)
		},
	}
}

// GetObject satisfies resource.crd interface (and retrieve.Retriever).
func (a *AutoScalerCRD) GetObject() runtime.Object {
	return &autoscalerdukev1alpha1.AutoScaler{}
}
