package k8s

import (
	apiextensionscli "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"

	pipelineclientset "github.com/marjoram/pipeline-operator/client/k8s/clientset/versioned"
	"github.com/marjoram/pipeline-operator/log"
)

// Service is the K8s service entrypoint.
type Services interface {
	CRD
	ConfigMap
	Pod
	PodDisruptionBudget
	RedisFailover
	Service
	RBAC
	Deployment
	StatefulSet
}

type services struct {
	CRD
	ConfigMap
	Pod
	PodDisruptionBudget
	RedisFailover
	Service
	RBAC
	Deployment
	StatefulSet
}

// New returns a new Kubernetes service.
func New(kubecli kubernetes.Interface, crdcli pipelineclientset.Interface, apiextcli apiextensionscli.Interface, logger log.Logger) Services {
	return &services{
		CRD:                 NewCRDService(apiextcli, logger),
		ConfigMap:           NewConfigMapService(kubecli, logger),
		Pod:                 NewPodService(kubecli, logger),
		PodDisruptionBudget: NewPodDisruptionBudgetService(kubecli, logger),
		RedisFailover:       NewRedisFailoverService(crdcli, logger),
		Service:             NewServiceService(kubecli, logger),
		RBAC:                NewRBACService(kubecli, logger),
		Deployment:          NewDeploymentService(kubecli, logger),
		StatefulSet:         NewStatefulSetService(kubecli, logger),
	}
}
