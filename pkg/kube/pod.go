package kube

func GeneratePodManifest() (string, error) {
    manifest := `
---
apiVersion: v1
kind: Pod
metadata:
  name: pod-name
  namespace: your-ns
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
`
    return manifest, nil
}
