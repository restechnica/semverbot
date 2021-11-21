package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/restechnica/semverbot/pkg/core"
)

// NewUpdateVersionCommand creates a new update version command.
// Returns the new spf13/cobra command.
func NewUpdateVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "version",
		RunE: UpdateVersionCommandRunE,
	}

	return command
}

// UpdateVersionCommandRunE runs the commands.
// Returns an error if it fails.
func UpdateVersionCommandRunE(cmd *cobra.Command, args []string) (err error) {
	if err = core.UpdateVersion(); err != nil {
		return fmt.Errorf("something went wrong while updating the version")
	}

	fmt.Println("successfully fetched the latest git tags")

	return err
}
