package main

import (
	"github.com/restechnica/semverbot/pkg/cli/exec"
)

// main bootstraps the `sbot` CLI app.
func main() {
	_ = exec.Run()
}
