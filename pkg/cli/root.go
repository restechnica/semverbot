package cli

import (
	"github.com/restechnica/semverbot/internal/config"
	"github.com/spf13/cobra"
)

const DefaultVersion = "v0.0.0"

func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:               "sbot",
		PersistentPreRunE: RootCommandPersistentPreRunE,
	}

	command.PersistentFlags().StringVarP(&ConfigFlag, "config", "c", "", "sbot config")

	command.AddCommand(NewGetCommand())
	command.AddCommand(NewPredictCommand())

	return command
}

func RootCommandPersistentPreRunE(cmd *cobra.Command, args []string) (err error) {
	Config = LoadDefaultConfig()

	if ConfigFlag != "" {
		Config, err = LoadConfig(ConfigFlag, Config)
	}

	return err
}

func LoadConfig(path string, target config.Root) (result config.Root, err error) {
	var loader = config.NewYAMLLoader()
	return loader.Overload(path, target)
}

func LoadDefaultConfig() config.Root {
	return config.NewRoot()
}
