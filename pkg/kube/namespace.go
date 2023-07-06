package kube

func GenerateNamespaceManifest() (string, error) {
    manifest := `
---
apiVersion: v1
kind: Namespace
metadata:
  name:  namespace-name
`
    return manifest, nil
}
