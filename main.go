package main

import (
	"log"

	"github.com/hsmtkk/urban-octo-garbanzo/command"
)

func main() {
	cmd := command.Command
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
