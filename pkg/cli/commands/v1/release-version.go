package v1

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/cli"
	"github.com/restechnica/semverbot/pkg/core"
)

// NewReleaseVersionCommand creates a new release version command.
// Returns the new spf13/cobra command.
func NewReleaseVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:     "version",
		PreRunE: ReleaseVersionCommandPreRunE,
		RunE:    ReleaseVersionCommandRunE,
	}

	command.Flags().StringVarP(&cli.ModeFlag, "mode", "m", "", "sbot mode")

	return command
}

// ReleaseVersionCommandPreRunE runs before the command runs.
// Returns an error if it fails.
func ReleaseVersionCommandPreRunE(cmd *cobra.Command, args []string) (err error) {
	return viper.BindPFlag(cli.ModeConfigKey, cmd.Flags().Lookup("mode"))
}

// ReleaseVersionCommandRunE runs the command.
// Returns an error if the command fails.
func ReleaseVersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
	log.Debug().Str("command", "release version").Msg("starting run...")

	var predictOptions = &core.PredictVersionOptions{
		DefaultVersion:      cli.DefaultVersion,
		GitBranchDelimiters: viper.GetString(cli.ModesGitBranchDelimitersConfigKey),
		GitCommitDelimiters: viper.GetString(cli.ModesGitCommitDelimitersConfigKey),
		Mode:                viper.GetString(cli.ModeConfigKey),
		SemverMap:           viper.GetStringMapStringSlice(cli.SemverMapConfigKey),
	}

	var releaseOptions = &core.ReleaseVersionOptions{
		GitTagsPrefix: viper.GetString(cli.GitTagsPrefixConfigKey),
	}

	log.Debug().
		Str("default", predictOptions.DefaultVersion).
		Str("mode", predictOptions.Mode).
		Str("prefix", releaseOptions.GitTagsPrefix).
		Msg("options")

	if err = core.ReleaseVersion(predictOptions, releaseOptions); err != nil {
		err = cli.NewCommandError(err)
	}

	return err
}
