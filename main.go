package main

import "github.com/restechnica/semverbot/cmd"

// main bootstraps the `sbot` CLI app.
func main() {
	var app = cmd.NewApp()
	_ = app.Execute()
}
