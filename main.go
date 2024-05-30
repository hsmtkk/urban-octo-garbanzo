package main

import (
	"log"

	"github.com/hsmtkk/wheat-future-sim/command"
)

func main() {
	cmd := command.Command
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
