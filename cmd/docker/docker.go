package docker

import (
	"context"
	"errors"
	"fmt"
	dockerType "github.com/docker/docker/api/types"
	docker "github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var option, service, deploymentDir string

func init() {
	rootDir, _ := os.Getwd()
	deploymentDir = rootDir + "/deployments"

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
		var client *docker.Client
		// Flag parser
		option, err = cmd.Flags().GetString("option")
		service, err = cmd.Flags().GetString("service")
		if err != nil {
			return err
		}

		// Docker API Client
		client, err = docker.NewClientWithOpts(docker.FromEnv)
		if err != nil {
			return err
		}
		// Make Docker API client to negotiate API Version due to supported API version Range
		client.NegotiateAPIVersion(context.Background())

		switch option {
		case "build":
			err = buildProjectImage(client)
		case "up":
		case "down":
		case "check":
		default:
			err = errors.New("Unsupported Option")
		}

		if err != nil {
			fmt.Println(service)
			return err
		}
		return nil
	},
}

func buildProjectImage(client *docker.Client) error {
	var err error
	var dockerFileReader *os.File
	var dockerResponse dockerType.ImageBuildResponse
	imageBuildOptions := dockerType.ImageBuildOptions{
		//Tags:       []string{"judgeengine"},
		NoCache:    true,
		Dockerfile: "Dockerfile",
	}

	dockerFileReader, err = os.Open(deploymentDir + "/Dockerfile")
	if err != nil {
		return err
	}
	dockerResponse, err = client.ImageBuild(context.Background(), dockerFileReader, imageBuildOptions)
	fmt.Println(dockerResponse)
	if err != nil {
		return err
	}
	defer dockerResponse.Body.Close()
	_, err = io.Copy(os.Stdout, dockerResponse.Body)
	if err != nil {
		return err
	}

	return nil
}
