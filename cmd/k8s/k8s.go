package k8s

import (
	"github.com/spf13/cobra"
)

func NewK8SCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "k8s",
		Short: "Generate Kubernetes manifest templates",
		Long: `Generate Kubernetes manifests on the given parameters.

This command generates Kubernetes manifests for various Kubernetes Objects.
You can pass the type of object you want to create.

Available Objects: configmap, deployment, ingress, namespace, pvc, service,
statefulset

Example usage:
  gockerize k8s service
`,
	}

	cmd.AddCommand(NewK8SServiceCommand())
	cmd.AddCommand(NewK8SDeploymentCommand())
	cmd.AddCommand(NewK8SNamespaceCommand())
	cmd.AddCommand(NewK8SIngressCommand())
	cmd.AddCommand(NewK8SConfigmapCommand())
	cmd.AddCommand(NewK8SPvcCommand())
	cmd.AddCommand(NewK8SStatefulsetCommand())
	// Add more Kubernetes object commands here

	return cmd
}
