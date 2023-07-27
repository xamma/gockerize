package k8s

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"github.com/xamma/gockerize/pkg/kube"
)

func NewK8SPodCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pod",
		Short: "Generate Kubernetes pod manifest template",
		Run: func(cmd *cobra.Command, args []string) {

			workMsg := color.New(color.FgYellow).Sprintf("Generating Kubernetes pod manifest template...")
			color.New(color.Bold).Println(workMsg)
			// use kube.GenerateServiceManifest from pkg/kube
			manifest, err := kube.GeneratePodManifest()
			if err != nil {
				log.Fatalf("Failed to generate Kubernetes pod manifest: %v", err)
			}
			// Save the manifest to a file
			filename := "pod.yaml"
			err = kube.SaveManifestToFile(manifest, filename)
			if err != nil {
				log.Fatalf("Failed to save Kubernetes pod manifest to file: %v", err)
			}

			succMsg := color.New(color.FgGreen).Sprintf("Kubernetes pod manifest template generated and saved to %s\n", filename)
			color.New(color.Bold).Println(succMsg)
		},
	}

	// command-specific flags go here (if needed)

	return cmd
}
