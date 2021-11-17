package commands

import (
	"fmt"
	"github.com/restechnica/semverbot/pkg/core"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/cli"
)

// NewPredictVersionCommand creates a new predict version command.
// returns the new spf13/cobra command.
func NewPredictVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:     "version",
		PreRunE: PredictVersionCommandPreRunE,
		RunE:    PredictVersionCommandRunE,
	}

	command.Flags().StringVarP(&cli.ModeFlag, "mode", "m", "auto", "sbot mode")

	return command
}

// PredictVersionCommandPreRunE runs before the command runs.
// returns an error if it fails.
func PredictVersionCommandPreRunE(cmd *cobra.Command, args []string) (err error) {
	return viper.BindPFlag(cli.SemverModeConfigKey, cmd.Flags().Lookup("mode"))
}

// PredictVersionCommandRunE runs the command.
// returns an error if the command fails.
func PredictVersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
	var options = &core.PredictVersionOptions{
		DefaultVersion:  cli.DefaultVersion,
		SemverDetection: viper.GetStringMapStringSlice(cli.SemverDetectionConfigKey),
		SemverMode:      viper.GetString(cli.SemverModeConfigKey),
	}

	var version string

	if version, err = core.PredictVersion(options); err == nil {
		fmt.Println(version)
	}

	return err
}
