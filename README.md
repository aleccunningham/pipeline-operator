# duked-crds
In-cluster, docker driven CI/CD pipeline for kubernetes:

```yaml
apiVersion: duked/v1alpha1
kind: Pipeline
metadata:
  name: project-pipeline
  labels:
    operator: pipeline-operator
spec:
  selector:
    app: {{ .Values.pipeline.proj }}
  reload: true
  workers:
    minWorkers: 1
    maxWorkers: 5
    resources:
      limits:
        cpu: 100m
        memory: 50Mi
  steps:
  - name: build
    volumeMounts:
    - name: build-scripts
      hostPath: /workspace/build  
  - name: test
    image: gcr.io/$PROJECT_ID/$image
    command: pytest -vv -cov
    artifacts:
    - name: coverage-reports
      dir: /workspace/cov
  volumes:
  - name: build-scripts
    gitURL: gist.github.com/alkdfaf.git
  - name: test-scripts
    mountPath: /workspace
  notify:
    when:
    - on_success
    - on_failure
    where:
      slack:
        channel:
        - general
        - example
        token:
        - name: slack-token
          hostPath: /workspace/config
      email:
        from_address: technology@company.com
        to_address:
        - dev1@gmail.com
        - dev2@gmail.com
```
