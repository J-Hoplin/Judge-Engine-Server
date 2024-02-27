package server

import (
	"github.com/spf13/cobra"
	"judge-engine/internal"
	"judge-engine/tools"
	"log"
)

var watch bool
var port int

func init() {
	ServerRootCmd.Flags().IntVarP(&port, "port", "p", 8080, "Set application port")
}

var ServerRootCmd = &cobra.Command{
	Use:   "run",
	Short: "Run application bootstrap function",
	Long:  "Set -w option to execute appplication server with dev mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		// Sync ent orm schema file
		log.Println("Sync ent ORM schema file...")
		var syncORMCommandString = []string{"go", "generate", "./ent"}
		var command = tools.ShellCommand{
			Command: syncORMCommandString,
			Visible: false,
		}
		if err = command.CommandBuilder().Run(); err != nil {
			return err
		}

		// Run server application
		err = internal.ApplicationBootstrap(port)
		return err
	},
}
