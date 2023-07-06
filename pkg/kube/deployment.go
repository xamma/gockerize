package kube

func GenerateDeploymentManifest() (string, error) {
    manifest := `
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: depl-name
  namespace: your-ns
spec:
  replicas: 3
  selector:
    matchLabels:
      app: label
  template:
    metadata:
      labels:
        app: label
    spec:
      containers:
      - name: container-name
        image: registry.io/image-name:latest
        imagePullPolicy: Always
        envFrom:
        - configMapRef:
            name: my-configmap
        ports:
        - containerPort: 8000
          name: port-name
      restartPolicy: Always
`
    return manifest, nil
}
