package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/internal"
)

// NewVersionCommand creates a new version command.
// It prints the version of the gitsync CLI.
// Returns the new version command.
func NewVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "version",
		Run: VersionCommandRun,
	}

	return command
}

// VersionCommandRun runs the command.
// TODO this command is not working as intended, still figuring out how to
// 		elegantly pass on build arguments, not so nice with makefile
func VersionCommandRun(cmd *cobra.Command, args []string) {
	fmt.Printf("sbot-cli %s go%s %s/%s\n", internal.Version, internal.GoVersion, internal.Os, internal.Arch)
}
