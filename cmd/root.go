package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/xamma/gockerize/cmd/k8s"
)

var rootCmd = &cobra.Command{
	Use:   "gockerize",
	Short: "A CLI tool to generate Dockerfiles for Go projects as well as K8s templates.",
	Long: `gockerize is a command-line tool that simplifies the process of generating Dockerfiles for Go projects.
It provides a convenient way to create Dockerfiles by specifying the project path, base images, build settings, and runtime configurations. 
Also, it allows you to create Kubernetes-templates for quick usage. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		welcomeSmiley := color.New(color.FgMagenta).Sprintf("ヽ༼ ಠ_ಠ༽ﾉ")
		color.New(color.Bold).Println(welcomeSmiley)
		fmt.Println()
		welcomeMsg := color.New(color.FgMagenta).Sprintf("GOCKERIZE")
		color.New(color.Bold).Println(welcomeMsg)
		fmt.Println()
		dockerMsg := color.New(color.FgCyan).Sprintf("To generate a Dockerfile for your Go project, use the 'generate' subcommand.")
		color.New(color.Bold).Println(dockerMsg)
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("  gockerize generate --path /path/to/project --build-base golang:1.16-alpine --run-base alpine:latest --build-output myapp --expose-port 8080 --cmd \"./myapp\"")
		fmt.Println()
		fmt.Println("For more information, use the 'generate --help' command.")
		fmt.Println()
		kubeMsg := color.New(color.FgCyan).Sprintf("To generate a template for a K8s object, use the 'k8s <objectname>' subcommand.")
		color.New(color.Bold).Println(kubeMsg)
		fmt.Println()
		discMsg := color.New(color.FgBlue).Sprintf("(c) 2023 Written by Max Bickel")
		color.New(color.FgMagenta).Println(discMsg)
	},
}

func init() {
	rootCmd.AddCommand(NewGenerateCommand())
	rootCmd.AddCommand(k8s.NewK8SCommand())
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
