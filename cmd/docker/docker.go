package docker

import (
	"errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

// Package level flag variable
var service, deploymentDir string

func init() {
	dir, _ := os.Getwd()
	deploymentDir = dir + "/deployments"

	// Global Flags

	// Local Flags
	DockerRootCmd.PersistentFlags().StringVarP(&service, "service", "s", "", "Specify service")
	DockerRootCmd.AddCommand(upCommand, downCommand, listCommand, buildCommand)
}

var DockerRootCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker related commands",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Subcommand not found")
	},
}

func getServiceList() ([]string, error) {
	type DockerComposeStruct struct {
		Version  string         `yaml:"version"`
		Services map[string]any `yaml:"services"`
	}
	var err error
	var yamlFile []byte
	var services = make([]string, 0)
	var yamlStruct = new(DockerComposeStruct)

	// Unmarshal docker compose file
	var composeFileDirectory = deploymentDir + "/docker-compose.yaml"
	yamlFile, err = os.ReadFile(composeFileDirectory)
	if err != nil {
		return nil, err
	}

	// Unmarshal yaml
	err = yaml.Unmarshal(yamlFile, &yamlStruct)
	if err != nil {
		return nil, err
	}

	for k, _ := range yamlStruct.Services {
		services = append(services, k)
	}
	return services, nil
}
