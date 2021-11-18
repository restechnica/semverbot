package main

import (
	"github.com/restechnica/semverbot/pkg/cli/commands"
)

// main bootstraps the `sbot` CLI app.
func main() {
	var cmd = commands.NewRootCommand()
	_ = cmd.Execute()
}
