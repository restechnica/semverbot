package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/internal/semver"
	"github.com/restechnica/semverbot/pkg/api"
	"github.com/restechnica/semverbot/pkg/cli"
)

func NewReleaseVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:     "version",
		PreRunE: ReleaseVersionCommandPreRunE,
		RunE:    ReleaseVersionCommandRunE,
	}

	command.Flags().StringVarP(&cli.ModeFlag, "mode", "m", "", "sbot mode")

	return command
}

func ReleaseVersionCommandPreRunE(cmd *cobra.Command, args []string) (err error) {
	return viper.BindPFlag("semver.mode", cmd.Flags().Lookup("mode"))
}

func ReleaseVersionCommandRunE(cmd *cobra.Command, args []string) error {
	var versionAPI = api.NewVersionAPI()
	var version = versionAPI.GetVersionOrDefault(cli.DefaultVersion)

	var mode = viper.GetString("semver.mode")
	var modeDetectionMap = viper.GetStringMapStringSlice("semver.modes.detection")
	var modeDetector = semver.NewModeDetector(modeDetectionMap)

	var semverModeAPI = api.NewSemverModeAPI(modeDetector)
	var semverMode = semverModeAPI.SelectMode(mode)

	var incrementedVersion string
	var err error

	if incrementedVersion, err = semverMode.Increment(version); err != nil {
		return err
	}

	var gitAPI = api.NewGitAPI()
	err = gitAPI.CreateAnnotatedTag(incrementedVersion)

	return err
}
