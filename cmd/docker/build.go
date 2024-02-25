package docker

import (
	"github.com/spf13/cobra"
	"judge-engine/tools"
)

var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "Compose up",
	Long:  "Compose service up. You can also select individual service.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if err != nil {
			return err
		}
		return buildProjectImage()
	},
}

func buildProjectImage() error {
	var err error
	var dockerfileDirectory = deploymentDir + "/Dockerfile"
	var dockerImageName = "judge-engine"

	var command = tools.ShellCommand{
		Command: []string{"docker", "build", "-t", dockerImageName, "-f", dockerfileDirectory, "."},
		Visible: true,
	}
	if err = command.CommandBuilder().Run(); err != nil {
		return err
	}
	return nil
}
