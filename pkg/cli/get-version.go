package cli

import (
	"fmt"
	"strings"

	"github.com/restechnica/semverbot/internal/commands"

	"github.com/spf13/cobra"
)

func NewGetVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "version",
		Run: GetVersionCommandRun,
	}

	return command
}

func GetVersionCommandRun(cmd *cobra.Command, args []string) {
	var version = GetVersionOrDefault(DefaultVersion)
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
