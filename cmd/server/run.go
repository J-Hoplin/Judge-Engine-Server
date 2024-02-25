package server

import (
	"github.com/spf13/cobra"
	"judge-engine/internal"
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
		err = internal.ApplicationBootstrap(port)
		return err
	},
}
