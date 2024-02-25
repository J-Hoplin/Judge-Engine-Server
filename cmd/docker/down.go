package docker

import (
	"errors"
	"github.com/spf13/cobra"
	"judge-engine/tools"
	"slices"
	"strings"
)

var downCommand = &cobra.Command{
	Use:   "down",
	Short: "Compose up",
	Long:  "Compose service up. You can also select individual service.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if err != nil {
			return err
		}
		return composeDown(service)
	},
}

func composeDown(service string) error {
	var err error
	var availableServiceList []string

	availableServiceList, err = getServiceList()

	var basicCommand = []string{"docker", "compose", "--project-directory", deploymentDir, "down"}

	if service != "" && !slices.Contains(availableServiceList, service) {
		return errors.New("Invalid service. Service must be one of: " + strings.Join(availableServiceList, ", "))
	}
	basicCommand = append(basicCommand, service)

	var command = tools.ShellCommand{
		Command: basicCommand,
		Visible: true,
	}
	if err = command.CommandBuilder().Run(); err != nil {
		return err
	}

	return nil
}
