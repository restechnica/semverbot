package commands

import (
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "get",
	}

	command.AddCommand(NewGetVersionCommand())

	return command
}
