package commands

import (
	"errors"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/restechnica/semverbot/pkg/cli"
	v1 "github.com/restechnica/semverbot/pkg/cli/commands/v1"
	"github.com/restechnica/semverbot/pkg/ext/viperx"
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
	command.AddCommand(v1.NewVersionCommand())

	return command
}

// RootCommandPersistentPreRunE runs before the command and any subcommand runs.
// Returns an error if it failed.
func RootCommandPersistentPreRunE(cmd *cobra.Command, args []string) (err error) {
	// silence usage output on error because errors at this point are unrelated to CLI usage
	cmd.SilenceUsage = true

	ConfigureLogging()

	log.Debug().Str("command", "root").Msg("starting pre-run...")

	log.Debug().Msg("loading default config values...")

	LoadDefaultConfigValues()

	if err = LoadConfigFile(cmd); err != nil {
		return err
	}

	if err = LoadFlagsIntoConfig(cmd); err != nil {
		return err
	}

	log.Debug().Msg("configuring git...")

	if err = SetGitConfigIfConfigured(); err != nil {
		return err
	}

	// silence errors which at this point are unrelated to CLI (cobra/viper) errors
	cmd.SilenceErrors = false

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

// LoadConfigFile loads the SemverBot configuration file.
// If the config flag was used, it will try to load only that path.
// If the config flag was not used, multiple default config file paths will be tried.
// Returns no error if config files are not found, returns an error if it fails otherwise.
func LoadConfigFile(cmd *cobra.Command) (err error) {
	configFlag := cmd.Flag("config")

	if configFlag.Changed {
		if err = viperx.LoadConfig(cli.ConfigFlag); err != nil {
			if errors.As(err, &viper.ConfigFileNotFoundError{}) {
				log.Warn().Msgf("config file %s not found", cli.ConfigFlag)
				return nil
			}
		}

		return err
	}

	paths := append([]string{cli.DefaultConfigFilePath}, cli.DefaultAdditionalConfigFilePaths...)

	for _, path := range paths {
		if err = viperx.LoadConfig(path); err == nil {
			return err
		}

		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return err
		}

		err = nil
		log.Warn().Msgf("config file %s not found", path)
	}

	return err
}

// LoadDefaultConfigValues loads the default SemverBot config.
func LoadDefaultConfigValues() {
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
