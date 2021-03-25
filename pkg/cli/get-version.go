package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewGetVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "version",
		RunE: RunGetVersionCommand,
	}

	return command
}

func RunGetVersionCommand(cmd *cobra.Command, args []string) (err error) {
	return GetVersion()
}

func GetVersion() (err error) {
	fmt.Println("v1.0.0")
	return
}
