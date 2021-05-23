package commands

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/api"
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
	command.AddCommand(NewUpdateCommand())

	return command
}

func RootCommandPersistentPreRunE(cmd *cobra.Command, args []string) (err error) {
	LoadDefaultConfig()

	if err = LoadConfig(); err != nil {
		return err
	}

	if err = LoadFlags(cmd); err != nil {
		return err
	}

	if err = SetGitConfigIfConfigured(); err != nil {
		return err
	}

	return err
}

func LoadConfig() (err error) {
	if cli.ConfigFlag != "" {
		viper.SetConfigFile(cli.ConfigFlag)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".semverbot")
		viper.SetConfigType("toml")
	}

	if err = viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			err = nil
		}
	}

	return err
}

func LoadDefaultConfig() {
	viper.SetDefault(cli.GitTagsPrefixConfigKey, "v")
	viper.SetDefault(cli.SemverDetectionConfigKey, map[string][]string{})
	viper.SetDefault(cli.SemverModeConfigKey, "auto")
}

func LoadFlags(cmd *cobra.Command) (err error) {
	return err
	//return viper.BindPFlag("git.tags.fetch", cmd.Flags().Lookup("fetch"))
}

func SetGitConfigIfConfigured() (err error) {
	var gitAPI = api.NewGitAPI()

	if viper.IsSet(cli.GitConfigEmailConfigKey) {
		var email = viper.GetString(cli.GitConfigEmailConfigKey)

		if err = gitAPI.SetConfigIfNotSet("user.email", email); err != nil {
			return err
		}
	}

	if viper.IsSet(cli.GitConfigNameConfigKey) {
		var name = viper.GetString(cli.GitConfigNameConfigKey)

		if err = gitAPI.SetConfigIfNotSet("user.name", name); err != nil {
			return err
		}
	}

	return err
}
