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

	command.PersistentFlags().BoolVarP(&cli.FetchFlag, "fetch", "f", false,
		`fetch all git tags before run (default "false")`)

	command.AddCommand(NewGetCommand())
	command.AddCommand(NewPredictCommand())
	command.AddCommand(NewReleaseCommand())

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

	if err = FetchGitTagsIfConfigured(); err != nil {
		return err
	}

	if err = SetGitConfigIfConfigured(); err != nil {
		return err
	}

	return err
}

func FetchGitTagsIfConfigured() (err error) {
	if viper.GetBool("git.tags.fetch") {
		err = api.NewGitAPI().FetchTags()
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
	viper.SetDefault("git.tags.fetch", false)
	viper.SetDefault("git.tags.prefix", "v")
	viper.SetDefault("semver.detection", map[string][]string{})
	viper.SetDefault("semver.mode", "auto")
}

func LoadFlags(cmd *cobra.Command) (err error) {
	return viper.BindPFlag("git.tags.fetch", cmd.Flags().Lookup("fetch"))
}

func SetGitConfigIfConfigured() (err error) {
	var gitAPI = api.NewGitAPI()

	if viper.IsSet("git.config.email") {
		var email = viper.GetString("git.config.email")

		if err = gitAPI.SetConfigIfNotSet("user.email", email); err != nil {
			return err
		}
	}

	if viper.IsSet("git.config.name") {
		var name = viper.GetString("git.config.name")

		if err = gitAPI.SetConfigIfNotSet("user.name", name); err != nil {
			return err
		}
	}

	return err
}
