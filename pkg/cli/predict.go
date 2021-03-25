package cli

import "github.com/spf13/cobra"

func NewPredictCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "predict",
	}

	command.AddCommand(NewPredictVersionCommand())

	return command
}
