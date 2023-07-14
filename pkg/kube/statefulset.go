package kube

func GenerateStatefulsetManifest() (string, error) {
    manifest := `
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: ss-name
  namespace: your-ns
spec:
  replicas: 1
  serviceName: service-name
  selector:
    matchLabels:
      app: label
  template:
    metadata:
      labels:
        app: label
    spec:
      containers:
        - name: minio
          image: bitnami/minio
          ports:
            - containerPort: 9000
              name: minio-port
          volumeMounts:
            - name: minio-vol
              mountPath: /data
  volumeClaimTemplates:
    - metadata:
        name: minio-vol
      spec:
        # storageClassName: local-path
        storageClassName: openebs-hostpath
        accessModes:
          - ReadWriteOnce
        resources:
          requests:
            storage: 1Gi
`
    return manifest, nil
}
