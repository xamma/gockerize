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
            - secretRef:
                name: my-secret
          env:
            - name: VAULT_SERVICE_NAME
              value: vault.ns.svc.cluster.local
            - name: GITHUB_TOKEN
              valueFrom:
                secretKeyRef:
                  name: github-token
                  key: token
          ports:
            - containerPort: 8000
              name: port-name
          volumeMounts:
            - name: volume-name
              mountPath: /app/path
      volumes:
        - name: volume-name
          persistentVolumeClaim:
            claimName: volume-claim-name
      restartPolicy: Always
`
    return manifest, nil
}
