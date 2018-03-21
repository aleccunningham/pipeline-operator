# Design

> An Operator is an application-specific controller that extends the Kubernetes API to create, configure and manage instances of complex stateful applications on behalf of a Kubernetes user. It builds upon the basic Kubernetes resource and controller concepts, but also includes domain or application-specific knowledge to automate common tasks better managed by computers.

```
    [ observe ] <--|--> [ trigger ]
         |         |
         |         |
         >         |
    [ analyze ]    |--> [ notify ]
         |         |
         |         |
         >         |
      [ act ] ---->|
```

The `pipeline-operator` ensures that a running instance of the `Pipeline` resource exists in the defined conditions at all times; which may mean spawning `Agent` resources, which are used to execute steps defined in your `Pipeline` CRD. 

A common use case could be described as follows. A certain `Deployment` would have a corresponding `Pipeline`, of which a series of steps, along with other resources and configurations, define what would be the equivilant of a `Jenkinsfile`, `.drone.yaml`, `.circle-ci.yaml`, etc. An `Agent` resource would also be created; however, it is not tied to a `Deployment` in the same way a `Pipeline` is - it instead acts as a worker to carry out `Pipeline` steps.

The operator runs a server that works in tanget with a git repository - it itself is pluggable, and may include a UI frontend. At its core, it implements interfaces for monitoring git and docker repositories, along with communicating with `Agent`'s for executing pipelines and reporting back to the operator, of which will then execute another step in the `Pipeline`.

Along with running an instance of the `Pipeline` server, the resource uses `spec.selector.pipeline` to update `Deployment`'s container image(s) when they source repository is updated (i.e. a change trigger from Github after a `git commit`). It is able to do so by specifying the `sourceRepository` spec in the `Pipeline` resource.

When an agent is requested from the operator, a pod is dynmically configured, based on the `Agent` spec, and executes the request by spawning containers for each `spec.pipeline.step`, with a persistent volume mounted on all pods at `/workspace` for pipeline artifacts, coverage reports, and user-specified directories. 
