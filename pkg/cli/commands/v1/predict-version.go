package v1

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/cli"
	"github.com/restechnica/semverbot/pkg/core"
)

// NewPredictVersionCommand creates a new predict version command.
// Returns the new spf13/cobra command.
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
// Returns an error if it fails.
func PredictVersionCommandPreRunE(cmd *cobra.Command, args []string) (err error) {
	return viper.BindPFlag(cli.ModeConfigKey, cmd.Flags().Lookup("mode"))
}

// PredictVersionCommandRunE runs the command.
// Returns an error if the command fails.
func PredictVersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
	log.Debug().Str("command", "v1.predict-version").Msg("starting run...")

	var options = &core.PredictVersionOptions{
		DefaultVersion:      cli.DefaultVersion,
		GitBranchDelimiters: viper.GetString(cli.ModesGitBranchDelimitersConfigKey),
		GitCommitDelimiters: viper.GetString(cli.ModesGitCommitDelimitersConfigKey),
		Mode:                viper.GetString(cli.ModeConfigKey),
		SemverMap:           viper.GetStringMapStringSlice(cli.SemverMapConfigKey),
	}

	log.Debug().
		Str("default", options.DefaultVersion).
		Str("mode", options.Mode).
		Msg("options")

	var version string

	if version, err = core.PredictVersion(options); err != nil {
		err = cli.NewCommandError(err)
	} else {
		fmt.Println(version)
	}

	return err
}
