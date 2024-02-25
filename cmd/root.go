package cmd

import (
	"github.com/spf13/cobra"
	"judge-engine/cmd/docker"
	"os"
)

// Root command
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Online Judge System Engine",
	Long:  "Engine API of Online Judge System",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Enroll Sub commands and flags
func init() {
	rootCmd.AddCommand(docker.DockerRootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
