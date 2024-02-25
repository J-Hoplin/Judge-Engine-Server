package docker

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"judge-engine/tools"
	"os"
	"slices"
	"strconv"
	"strings"
)

var option, service, deploymentDir string

func init() {
	dir, _ := os.Getwd()
	deploymentDir = dir + "/deployments"

	// Global Flags

	// Local Flags
	DockerRootCmd.Flags().StringVarP(&option, "option", "o", "", "Option of exeuting docker")
	DockerRootCmd.Flags().StringVarP(&service, "service", "s", "", "Specify service")
}

var DockerRootCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker related commands",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var option, service string
		// Flag parser
		option, err = cmd.Flags().GetString("option")
		service, err = cmd.Flags().GetString("service")
		if err != nil {
			return err
		}

		switch option {
		case "list":
			err = listServices()
		case "build":
			err = buildProjectImage()
		case "up":
			err = composeUp(service)
		case "down":
			err = composeDown(service)
		default:
			err = errors.New("Unsupported Option")
		}

		if err != nil {
			return err
		}
		return nil
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

func listServices() error {
	var err error
	var serviceNames []string
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	serviceNames, err = getServiceList()
	if err != nil {
		return err
	}

	for index, value := range serviceNames {
		fmt.Fprintln(writer, strconv.Itoa(index+1)+". "+value)
	}
	return nil
}

func composeUp(service string) error {
	var err error
	var availableServiceList []string

	availableServiceList, err = getServiceList()

	var basicCommand = []string{"docker", "compose", "--project-directory", deploymentDir, "up", "-d"}

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
