package k8s

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"github.com/xamma/gockerize/pkg/kube"
)

func NewK8SIngressCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingress",
		Short: "Generate Kubernetes ingress manifest template",
		Run: func(cmd *cobra.Command, args []string) {

			workMsg := color.New(color.FgYellow).Sprintf("Generating Kubernetes ingress manifest template...")
			color.New(color.Bold).Println(workMsg)
			// use kube.GenerateServiceManifest from pkg/kube
			manifest, err := kube.GenerateIngressManifest()
			if err != nil {
				log.Fatalf("Failed to generate Kubernetes ingress manifest: %v", err)
			}
			// Save the manifest to a file
			filename := "ingress.yaml"
			err = kube.SaveManifestToFile(manifest, filename)
			if err != nil {
				log.Fatalf("Failed to save Kubernetes ingress manifest to file: %v", err)
			}

			succMsg := color.New(color.FgGreen).Sprintf("Kubernetes ingress manifest template generated and saved to %s\n", filename)
			color.New(color.Bold).Println(succMsg)
		},
	}

	// command-specific flags go here (if needed)

	return cmd
}
