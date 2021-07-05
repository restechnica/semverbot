package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/internal/semver"
	"github.com/restechnica/semverbot/pkg/api"
	"github.com/restechnica/semverbot/pkg/cli"
)

// NewReleaseVersionCommand creates a new release version command.
// returns the new spf13/cobra command.
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
// returns an error if it fails.
func ReleaseVersionCommandPreRunE(cmd *cobra.Command, args []string) (err error) {
	return viper.BindPFlag(cli.SemverModeConfigKey, cmd.Flags().Lookup("mode"))
}

// ReleaseVersionCommandRunE runs the command.
// returns an error if the command fails.
func ReleaseVersionCommandRunE(cmd *cobra.Command, args []string) error {
	var versionAPI = api.NewVersionAPI()
	var version = versionAPI.GetVersionOrDefault(cli.DefaultVersion)

	var mode = viper.GetString(cli.SemverModeConfigKey)
	var modeDetectionMap = viper.GetStringMapStringSlice(cli.SemverDetectionConfigKey)
	var modeDetector = semver.NewModeDetector(modeDetectionMap)

	var semverModeAPI = api.NewSemverModeAPI(modeDetector)
	var semverMode = semverModeAPI.SelectMode(mode)

	var incrementedVersion string
	var err error

	if incrementedVersion, err = semverMode.Increment(version); err != nil {
		return err
	}

	var gitTagPrefix = viper.GetString(cli.GitTagsPrefixConfigKey)
	incrementedVersion = fmt.Sprintf("%s%s", gitTagPrefix, incrementedVersion)

	var gitAPI = api.NewGitAPI()
	err = gitAPI.CreateAnnotatedTag(incrementedVersion)

	return err
}
