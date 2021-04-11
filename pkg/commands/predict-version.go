package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/internal/semver"
	"github.com/restechnica/semverbot/pkg/api"
	"github.com/restechnica/semverbot/pkg/cli"
)

func NewPredictVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:     "version",
		PreRunE: PredictVersionCommandPreRunE,
		RunE:    PredictVersionCommandRunE,
	}

	command.Flags().StringVarP(&cli.ModeFlag, "mode", "m", "auto", "sbot mode")

	return command
}

func PredictVersionCommandPreRunE(cmd *cobra.Command, args []string) (err error) {
	return viper.BindPFlag("semver.mode", cmd.Flags().Lookup("mode"))
}

func PredictVersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
	var versionAPI = api.NewVersionAPI()
	var version = versionAPI.GetVersionOrDefault(cli.DefaultVersion)

	var mode = viper.GetString("semver.mode")
	var modeDetectionMap = viper.GetStringMapStringSlice("semver.modes.detection")
	var modeDetector = semver.NewModeDetector(modeDetectionMap)

	var semverModeAPI = api.NewSemverModeAPI(modeDetector)
	var semverMode = semverModeAPI.SelectMode(mode)

	var incrementedVersion string

	if incrementedVersion, err = semverMode.Increment(version); err != nil {
		return err
	}

	fmt.Println(incrementedVersion)

	return err
}
