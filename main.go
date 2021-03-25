package main

import (
	"log"

	"github.com/restechnica/semverbot/cmd"
)

func main() {
	var app = cmd.NewApp()

	if err := app.Execute(); err != nil {
		log.Fatal(err)
	}
}
