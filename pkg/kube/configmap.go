package kube

func GenerateConfigmapManifest() (string, error) {
    manifest := `
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: configmap-name
  namespace: your-ns
data:
  KEY_LIST1: >-
    val1,
    val2,
    val3,
    val4,
    val5,
    val6
  KEY1: "val1"
  KEY2: "val2"
  KEY3: "val3"
  EXAMPLE_SERVICE_HOST: http://servicename:8080/api/v1/endpoint
`
    return manifest, nil
}
