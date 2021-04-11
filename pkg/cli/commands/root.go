package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/internal/semver"
	"github.com/restechnica/semverbot/pkg/cli"
)

func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:               "sbot",
		PersistentPreRunE: RootCommandPersistentPreRunE,
	}

	command.PersistentFlags().StringVarP(&cli.ConfigFlag, "config", "c", "",
		`sbot config (default ".semverbot.toml")`)

	command.AddCommand(NewGetCommand())
	command.AddCommand(NewPredictCommand())
	command.AddCommand(NewReleaseCommand())

	return command
}

func RootCommandPersistentPreRunE(cmd *cobra.Command, args []string) (err error) {
	LoadDefaultConfig()
	return LoadConfig()
}

func LoadConfig() (err error) {
	if cli.ConfigFlag != "" {
		viper.SetConfigFile(cli.ConfigFlag)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".semverbot")
		viper.SetConfigType("toml")
	}

	return viper.ReadInConfig()
}

func LoadDefaultConfig() {
	viper.SetDefault("semver.matchers", []semver.Mode{})
	viper.SetDefault("semver.mode", "auto")
}
