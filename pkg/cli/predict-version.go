package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewPredictVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "version",
		Run: PredictVersionCommandRun,
	}

	command.PersistentFlags().StringVarP(&ModeFlag, "mode", "m", "auto", "sbot mode")

	return command
}

func PredictVersionCommandRun(cmd *cobra.Command, args []string) {
	var version = PredictVersion()
	fmt.Println(version)
}

func PredictVersion() (version string) {
	//var mnger = semver.NewManager()
	version = GetVersionOrDefault(DefaultVersion)

	return
}
