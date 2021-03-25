package main

import (
	"github.com/restechnica/semverbot/cmd"
)

func main() {
	var app = cmd.NewApp()
	_ = app.Execute()
}
