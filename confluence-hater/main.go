package main

import (
	"os"

	"github.com/lovemew67/project-misc/confluence-hater/command"
)

func main() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
