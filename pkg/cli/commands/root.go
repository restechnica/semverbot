package commands

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/cli"
	"github.com/restechnica/semverbot/pkg/git"
)

// NewRootCommand creates a new root command.
// Returns the new spf13/cobra command.
func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:               "sbot",
		PersistentPreRunE: RootCommandPersistentPreRunE,
	}

	command.PersistentFlags().StringVarP(&cli.ConfigFlag, "config", "c", "",
		fmt.Sprintf(`configures which config file to use (default "%s")`, cli.DefaultConfigFilePath))

	command.AddCommand(NewGetCommand())
	command.AddCommand(NewInitCommand())
	command.AddCommand(NewPredictCommand())
	command.AddCommand(NewPushCommand())
	command.AddCommand(NewReleaseCommand())
	command.AddCommand(NewUpdateCommand())

	return command
}

// RootCommandPersistentPreRunE runs before the command and any subcommand runs.
// Returns an error if it failed.
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

// LoadConfig loads the semverbot configuration file.
// Returns an error if it fails.
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

// LoadDefaultConfig loads the default semverbot config.
func LoadDefaultConfig() {
	viper.SetDefault(cli.GitTagsPrefixConfigKey, "v")
	viper.SetDefault(cli.SemverMatchConfigKey, map[string][]string{})
	viper.SetDefault(cli.SemverModeConfigKey, "auto")
}

// LoadFlags loads root command flags.
// Returns an error if it fails.
func LoadFlags(cmd *cobra.Command) (err error) {
	return err
	//return viper.BindPFlag("git.tags.fetch", cmd.Flags().Lookup("fetch")) -- example on how to load flags
}

// SetGitConfigIfConfigured Sets the git config only when the semverbot config exists and
// the git config does not exist.
// Returns an error if it fails.
func SetGitConfigIfConfigured() (err error) {
	var gitAPI = git.NewAPI()

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
