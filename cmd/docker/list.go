package docker

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Compose up",
	Long:  "Compose service up. You can also select individual service.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if err != nil {
			return err
		}
		return listServices()
	},
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
