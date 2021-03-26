package commands

import (
	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/internal/config"
	"github.com/restechnica/semverbot/pkg/cli"
)

func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:               "sbot",
		PersistentPreRunE: RootCommandPersistentPreRunE,
	}

	command.PersistentFlags().StringVarP(&cli.ConfigFlag, "config", "c", "", "sbot config")

	command.AddCommand(NewGetCommand())
	command.AddCommand(NewPredictCommand())

	return command
}

func RootCommandPersistentPreRunE(cmd *cobra.Command, args []string) (err error) {
	cli.Config = LoadDefaultConfig()

	if cli.ConfigFlag != "" {
		cli.Config, err = LoadConfig(cli.ConfigFlag, cli.Config)
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
