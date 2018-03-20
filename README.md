# Pipeline Operator

**Project status: *alpha*** Nearly all planned features are incomplete. The API, spec, status and other user facing objects will introduce breaking changes.

The Pipeline Operator for Kubernetes provides easy continuous delivery for Kubernetes deployments and management of Pipeline instances.

Once installed, the Pipeline Operator provides the following features:

* **Create/Destroy**: Easily launch a Pipeline instance for your Kubernetes namespace,
  a specific application or team easily using the Operator.

* **Simple Configuration**: Configure the fundamentals of Pipeline like versions, persistence,
  retention policies, and replicas from a native Kubernetes resource.

* **Target Services via Labels**: Automatically generate build trigger configurations based
  on familiar Kubernetes label queries; no need to setup any external CI platform.

The current project roadmap [can be found here](./ROADMAP.md).

## Prerequisites

The Pipeline Operator requires a Kubernetes cluster of version `>=1.9.0`. If you are just starting out with the
Pipeline Operator, it is highly recommended to use the latest version.

## CustomResourceDefinitions

The Operator acts on the following [custom resource definitions (CRDs)](https://kubernetes.io/docs/tasks/access-kubernetes-api/extend-api-custom-resource-definitions/):

* **`Pipeline`**, which defines a desired Pipeline deployment and it's related deployment.
  The Operator ensures at all times that a deployment matching the resource definition is running.

* **`Agent`**, which declaratively specifies in-cluster pipeline-runners
  The Operator automatically generates Pipeline configuration for the agent based on the definition.
  
* **`AutoScaler`**, Automatically adds or removes instances based on build volume 
  WIP Move `AutoScaler` into `Agent` operator

To learn more about the CRDs introduced by the Pipeline Operator have a look
at the [design doc](docs/design.md).

## Installation

Install the Operator inside a cluster by running the following command:

```
kubectl apply -f bundle.yaml
```

> Note: make sure to adapt the namespace in the ClusterRoleBinding if deploying in another namespace than the default namespace.

## Removal

To remove the operator and Prometheus, first delete any custom resources you created in each namespace. The
operator will automatically shut down and remove Prometheus and Alertmanager pods, and associated configmaps.

```
for n in $(kubectl get namespaces -o jsonpath={..metadata.name}); do
  kubectl delete --all --namespace=$n pipeline,agent
done
```

After a couple of minutes you can go ahead and remove the operator itself.

```
kubectl delete -f bundle.yaml
```

The operator automatically creates services in each namespace where you created a Prometheus or Alertmanager resources,
and defines three custom resource definitions. You can clean these up now.

```
for n in $(kubectl get namespaces -o jsonpath={..metadata.name}); do
  kubectl delete --ignore-not-found --namespace=$n service pipeline-operater agent-operater
done

kubectl delete --ignore-not-found customresourcedefinitions pipeline.duke.lol agent.duke.lol
```

## Development

### Prerequisites

- golang environment
- docker (used for creating container images, etc.)
- minikube (optional)

### `minikube`

Pass the following flags to `minikube start` to enable neccesary configuration changes:
```
#!/bin/bash
minikube delete; minikube start --kubernetes-version=v1.9.1 --memory=4096 --bootstrapper=kubeadm --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.address=0.0.0.0 --extra-config=controller-manager.address=0.0.0.0
```

When making changes to the API, ensure that you update the versioned clients via `make generate`.
