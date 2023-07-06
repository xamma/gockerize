package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

var rootCmd = &cobra.Command{
	Use:   "gockerize",
	Short: "A CLI tool to generate Dockerfiles for Go projects",
	Long: `gockerize is a command-line tool that simplifies the process of generating Dockerfiles for Go projects.
It provides a convenient way to create Dockerfiles by specifying the project path, base images, build settings, and runtime configurations. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		welcomeMsg := color.New(color.FgBlue).Sprintf("Welcome to gockerize!")
		color.New(color.Bold).Println(welcomeMsg)
		fmt.Println()
		fmt.Println("To generate a Dockerfile for your Go project, use the 'generate' subcommand.")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("  gockerize generate --path /path/to/project --build-base golang:1.16-alpine --run-base alpine:latest --build-output myapp --expose-port 8080 --cmd \"./myapp\"")
		fmt.Println()
		fmt.Println("For more information, use the 'generate --help' command.")
	},
}

func init() {
	rootCmd.AddCommand(NewGenerateCommand())
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
