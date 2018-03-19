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
  workers:
    minWorkers: 1
    maxWorkers: 5
    resources:
      limits:
        cpu: 100m
        memory: 50Mi
  jobs:
  - name: build
    image: gcr.io/$PROJECT_ID/$image
    imagePullPolicy: IfNotPresent
    volumeMounts:
    - name: build-scripts
      hostPath: /workspace/build  
  - name: test
    image: gcr.io/$PROJECT_ID/$image
    volumeMounts:
    - name: test-scripts
      hostPath: /workspace/test
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
      - technology@company.com
```
