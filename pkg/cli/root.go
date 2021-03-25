package cli

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "sbot",
	}

	command.PersistentFlags().StringVarP(&Config, "config", "c", "", "sbot config")

	command.AddCommand(NewGetCommand())
	command.AddCommand(NewPredictCommand())

	return command
}
