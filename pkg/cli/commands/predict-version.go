package commands

import (
	"fmt"

	"github.com/restechnica/semverbot/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/cli"
)

func NewPredictVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "version",
		RunE: PredictVersionCommandRunE,
	}

	command.PersistentFlags().StringVarP(&cli.ModeFlag, "mode", "m", "auto", "sbot mode")
	_ = viper.BindPFlag("semver.mode", command.PersistentFlags().Lookup("mode"))

	return command
}

func PredictVersionCommandRunE(cmd *cobra.Command, args []string) error {
	var versionAPI = api.NewVersionAPI()
	var version = versionAPI.GetVersionOrDefault(cli.DefaultVersion)

	var mode = viper.GetString("semver.mode")
	var gitCommitMatchers = viper.GetStringMapString("semver.matchers")
	var semverModeAPI = api.NewSemverModeAPI(nil, gitCommitMatchers)
	var semverMode = semverModeAPI.SelectMode(mode)

	var incrementedVersion, err = semverMode.Increment(version)
	fmt.Println(incrementedVersion)

	return err
}
