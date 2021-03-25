package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewPredictVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "version",
		RunE: RunPredictVersionCommand,
	}

	return command
}

func RunPredictVersionCommand(cmd *cobra.Command, args []string) (err error) {
	return PredictVersion()
}

func PredictVersion() (err error) {
	fmt.Println("v1.0.0")
	return
}
