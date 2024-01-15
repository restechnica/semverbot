package commands

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/cli"
	v1 "github.com/restechnica/semverbot/pkg/cli/commands/v1"
	"github.com/restechnica/semverbot/pkg/git"
	"github.com/restechnica/semverbot/pkg/semver"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
}

// NewRootCommand creates a new root command.
// Returns the new spf13/cobra command.
func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:               "sbot",
		PersistentPreRunE: RootCommandPersistentPreRunE,
	}

	command.PersistentFlags().StringVarP(&cli.ConfigFlag, "config", "c", cli.DefaultConfigFilePath, "configures which config file to use")

	command.PersistentFlags().BoolVarP(&cli.VerboseFlag, "verbose", "v", false, "increase log level verbosity to Info")
	command.PersistentFlags().BoolVarP(&cli.DebugFlag, "debug", "d", false, "increase log level verbosity to Debug")

	command.AddCommand(v1.NewV1Command())
	command.AddCommand(v1.NewGetCommand())
	command.AddCommand(v1.NewInitCommand())
	command.AddCommand(v1.NewPredictCommand())
	command.AddCommand(v1.NewPushCommand())
	command.AddCommand(v1.NewReleaseCommand())
	command.AddCommand(v1.NewUpdateCommand())
	command.AddCommand(NewVersionCommand())

	return command
}

// RootCommandPersistentPreRunE runs before the command and any subcommand runs.
// Returns an error if it failed.
func RootCommandPersistentPreRunE(cmd *cobra.Command, args []string) (err error) {
	// silence usage and errors because errors at this point are unrelated to CLI usage errors
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true

	ConfigureLogging()

	log.Debug().Str("command", "root").Msg("starting pre-run...")

	log.Debug().Msg("loading default config...")

	LoadDefaultConfig()

	if err = LoadConfig(); err != nil {
		return err
	}

	if err = LoadFlagsIntoConfig(cmd); err != nil {
		return err
	}

	log.Debug().Msg("configuring git...")

	if err = SetGitConfigIfConfigured(); err != nil {
		return err
	}

	return err
}

func ConfigureLogging() {
	SetLogLevel()
}

func SetLogLevel() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	if cli.VerboseFlag {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if cli.DebugFlag {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

// LoadConfig loads the SemverBot configuration file.
// Returns an error if it fails.
func LoadConfig() (err error) {
	viper.AddConfigPath(filepath.Dir(cli.ConfigFlag))
	viper.SetConfigName(strings.TrimSuffix(filepath.Base(cli.ConfigFlag), filepath.Ext(cli.ConfigFlag)))
	viper.SetConfigType(strings.Split(filepath.Ext(cli.ConfigFlag), ".")[1])

	log.Debug().Str("path", cli.ConfigFlag).Msg("loading config...")

	if err = viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			log.Warn().Msg("config file not found")
			return nil
		}
	}

	return err
}

// LoadDefaultConfig loads the default SemverBot config.
func LoadDefaultConfig() {
	viper.SetDefault(cli.GitTagsPrefixConfigKey, cli.DefaultGitTagsPrefix)
	viper.SetDefault(cli.ModeConfigKey, cli.DefaultMode)
	viper.SetDefault(cli.ModesGitBranchDelimitersConfigKey, cli.DefaultGitBranchDelimiters)
	viper.SetDefault(cli.ModesGitCommitDelimitersConfigKey, cli.DefaultGitCommitDelimiters)
	viper.SetDefault(cli.SemverMapConfigKey, semver.Map{})
}

// LoadFlagsIntoConfig loads root command flags.
// Returns an error if it fails.
func LoadFlagsIntoConfig(cmd *cobra.Command) (err error) {
	return err
	//return viper.BindPFlag("git.tags.fetch", cmd.Flags().Lookup("fetch")) -- example on how to load flags
}

// SetGitConfigIfConfigured Sets the git config only when the SemverBot config exists and the git config does not exist.
// Returns an error if it fails.
func SetGitConfigIfConfigured() (err error) {
	var gitAPI = git.NewCLI()
	var value string

	if viper.IsSet(cli.GitConfigEmailConfigKey) {
		var email = viper.GetString(cli.GitConfigEmailConfigKey)
		value = email

		if value, err = gitAPI.SetConfigIfNotSet("user.email", email); err != nil {
			return err
		}

	}

	log.Debug().Str("user.email", strings.Trim(value, "\n")).Msg("")

	if viper.IsSet(cli.GitConfigNameConfigKey) {
		var name = viper.GetString(cli.GitConfigNameConfigKey)
		value = name

		if value, err = gitAPI.SetConfigIfNotSet("user.name", name); err != nil {
			return err
		}

	}

	log.Debug().Str("user.name", strings.Trim(value, "\n")).Msg("")

	return err
}
