package cmd

import (
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/xamma/gockerize/pkg/dockerizer"
)

func NewGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate Dockerfile for a Go/Python project",
		Long: `Generate Dockerfile for a Go/Python project based on the given parameters.

This command generates a Dockerfile for a Go/Python project based on the provided options.
It uses multi-stage builds to create minimal Docker images.
Go is the default project-type if you dont pass values.

Example usage:
  gockerize generate --path /path/to/project --build-base golang:1.16-alpine --run-base alpine:latest --build-output myapp --expose-port 8080 --cmd "./myapp"
`,
		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("path")
			buildBaseImage, _ := cmd.Flags().GetString("build-base")
			runBaseImage, _ := cmd.Flags().GetString("run-base")
			buildOutput, _ := cmd.Flags().GetString("build-output")
			exposePort, _ := cmd.Flags().GetString("expose-port")
			cmdValue, _ := cmd.Flags().GetString("cmd")
			dockerfilePath, _ := cmd.Flags().GetString("dockerfile-path")
			projectType, _ := cmd.Flags().GetString("project-type")

			err := dockerizer.GenerateDockerfile(dockerfilePath, path, buildBaseImage, runBaseImage, buildOutput, exposePort, cmdValue, projectType)
			if err != nil {
				log.Fatalf("Failed to generate Dockerfile: %v", err)
			}

			// Output success message in green color
			successMsg := color.New(color.FgGreen).Sprintf("Dockerfile generated successfully.")
			color.New(color.Bold).Println(successMsg)
		},
	}

	cmd.Flags().String("path", ".", "Path to the Go project")
	cmd.Flags().String("build-base", "golang:1.20-alpine", "Base image for the build stage of the Dockerfile")
	cmd.Flags().String("run-base", "alpine:latest", "Base image for the run stage of the Dockerfile")
	cmd.Flags().String("build-output", "goapp", "Name of the build output binary")
	cmd.Flags().String("expose-port", "8080", "Port to expose in the Dockerfile")
	cmd.Flags().String("cmd", "./goapp", "Command to run in the Dockerfile CMD")
	cmd.Flags().String("dockerfile-path", "", "Path to the Dockerfile")
	cmd.Flags().String("project-type", "go", "Project Type: Go or Python")

	// cmd.MarkFlagRequired("path")

	return cmd
}
