package main

import (
	"os"

	"github.com/lovemew67/public-misc/golang-sample/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
