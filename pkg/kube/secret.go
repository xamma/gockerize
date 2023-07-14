package kube

func GenerateSecretManifest() (string, error) {
    manifest := `
---
apiVersion: v1
kind: Secret
metadata:
  name: secret-name
  namespace: your-ns
type: Opaque
data:
  token: VERYBASESECRET
  # echo -n "VERYSECRET" | base64
  # kubectl create secret docker-registry ghcr-secret --docker-server=https://ghcr.io --docker-username=<YOURNAME> --docker-password=XXX --docker-email=bla@test.com -n your-ns --dry-run=client -o yaml | kubectl apply -f -
`
    return manifest, nil
}