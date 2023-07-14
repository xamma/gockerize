package kube

func GenerateServiceManifest() (string, error) {
    manifest := `
---
apiVersion: v1
kind: Service
metadata:
  name: servicename
  namespace: your-ns
spec:
  selector:
    app: label
  type: ClusterIP
  ports:
    - name: portname
      protocol: TCP
      port: 9000
      targetPort: 9000
`
    return manifest, nil
}
