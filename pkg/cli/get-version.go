package cli

import (
	"fmt"
	"github.com/restechnica/semverbot/internal/commands"
	"strings"

	"github.com/spf13/cobra"
)

func NewGetVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "version",
		Run: RunGetVersionCommand,
	}

	return command
}

func RunGetVersionCommand(cmd *cobra.Command, args []string) {
	var version = GetVersionOrDefault("v0.0.0")
	fmt.Println(version)
}

func GetVersion() (version string, err error) {
	var cmder = commands.ExecCommander{}
	version, err = cmder.Output("git", "describe", "--tags")
	return strings.TrimSpace(version), err
}

func GetVersionOrDefault(defaultVersion string) (version string) {
	var err error

	if version, err = GetVersion(); err != nil {
		version = defaultVersion
	}

	return version
}
