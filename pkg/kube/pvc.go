package kube

func GeneratePvcManifest() (string, error) {
    manifest := `
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc-claim
  namespace: your-ns
  labels:
    app: label
spec:
  # storageClassName: longhorn
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
