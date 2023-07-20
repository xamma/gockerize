package kube

func GenerateServiceAccountManifest() (string, error) {
    manifest := `
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cluster-admin
  namespace: your-ns

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cluster-admin
rules:
  - apiGroups:
      - "*"
    resources:
      - "*"
    verbs:
      - "*"

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cluster-admin
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: cluster-admin
    namespace: your-ns
`
    return manifest, nil
}
