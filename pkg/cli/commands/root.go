package commands

import (
	"github.com/restechnica/semverbot/internal/semver"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/cli"
)

func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:               "sbot",
		PersistentPreRunE: RootCommandPersistentPreRunE,
	}

	command.PersistentFlags().StringVarP(&cli.ConfigFlag, "config", "c", "",
		`sbot config (default ".semverbot.yaml")`)

	command.AddCommand(NewGetCommand())
	command.AddCommand(NewPredictCommand())

	return command
}

func RootCommandPersistentPreRunE(cmd *cobra.Command, args []string) (err error) {
	SetConfigDefaults()
	err = LoadConfig()
	return err
}

func LoadConfig() (err error) {
	if cli.ConfigFlag != "" {
		viper.SetConfigFile(cli.ConfigFlag)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".semverbot")
		viper.SetConfigType("yaml")
	}

	return viper.ReadInConfig()
}

func SetConfigDefaults() {
	viper.SetDefault("semver.mode", "auto")
	viper.SetDefault("semver.matchers", []semver.Mode{})
}
