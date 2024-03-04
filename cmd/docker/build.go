package docker

import (
	"github.com/spf13/cobra"
	"judge-engine/tools"
	"log"
	"sync"
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
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	var mutex *sync.Mutex = new(sync.Mutex)
	var apiDockerfileDirectory = deploymentDir + "/ApiDockerfile"
	var isolateDockerfileDirectory = deploymentDir + "/IsolateDockerfile"
	var apiDockerImageName = "judge-api"
	var isolateDockerImageName = "judge-isolate"

	wg.Add(2)

	go func() {
		var innerErr error
		var command tools.ShellCommand
		log.Println("Build docker image for API server")
		command = tools.ShellCommand{
			Command: []string{"docker", "build", "-t", apiDockerImageName, "-f", apiDockerfileDirectory, "."},
			Visible: true,
		}
		if innerErr = command.CommandBuilder().Run(); innerErr != nil {
			mutex.Lock()
			err = innerErr
			mutex.Unlock()
		}
		wg.Done()
	}()

	go func() {
		var innerErr error
		var command tools.ShellCommand
		log.Println("Build docker image for Isolate Environment")
		command = tools.ShellCommand{
			Command: []string{"docker", "build", "-t", isolateDockerImageName, "-f", isolateDockerfileDirectory, "."},
			Visible: true,
		}
		if innerErr = command.CommandBuilder().Run(); innerErr != nil {
			mutex.Lock()
			err = innerErr
			mutex.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()

	return err
}
