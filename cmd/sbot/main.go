package main

import (
	"github.com/restechnica/semverbot/pkg/cli/commands"
)

// main bootstraps the `sbot` CLI app.
func main() {
	var cli = commands.NewRootCommand()
	_ = cli.Execute()
}
