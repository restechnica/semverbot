package cmd

import (
	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/pkg/cli"
)

func NewApp() *cobra.Command {
	return cli.NewRootCommand()
}
