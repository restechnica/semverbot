package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/api"
	"github.com/restechnica/semverbot/pkg/cli"
)

func NewPushVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "version",
		RunE: PushVersionCommandRunE,
	}

	return command
}

func PushVersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
	var versionAPI = api.NewVersionAPI()
	var version = versionAPI.GetVersionOrDefault(cli.DefaultVersion)

	var gitTagPrefix = viper.GetString(cli.GitTagsPrefixConfigKey)
	var prefixedVersion = fmt.Sprintf("%s%s", gitTagPrefix, version)

	var gitAPI = api.NewGitAPI()
	return gitAPI.PushTag(prefixedVersion)
}
