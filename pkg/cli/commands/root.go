package commands

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/cli"
	"github.com/restechnica/semverbot/pkg/cli/commands/v1"
	"github.com/restechnica/semverbot/pkg/git"
	"github.com/restechnica/semverbot/pkg/semver"
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

	command.AddCommand(v1.NewV1Command())
	command.AddCommand(v1.NewGetCommand())
	command.AddCommand(v1.NewInitCommand())
	command.AddCommand(v1.NewPredictCommand())
	command.AddCommand(v1.NewPushCommand())
	command.AddCommand(v1.NewReleaseCommand())
	command.AddCommand(v1.NewUpdateCommand())

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
	viper.SetDefault(cli.GitTagsPrefixConfigKey, cli.DefaultGitTagsPrefix)
	viper.SetDefault(cli.ModeConfigKey, cli.DefaultMode)
	viper.SetDefault(cli.ModesGitBranchDelimitersConfigKey, cli.DefaultGitBranchDelimiters)
	viper.SetDefault(cli.ModesGitCommitDelimitersConfigKey, cli.DefaultGitCommitDelimiters)
	viper.SetDefault(cli.SemverMapConfigKey, semver.Map{})
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
	var gitAPI = git.NewCLI()

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
