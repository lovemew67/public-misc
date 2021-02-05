package main

import (
	"os"

	"github.com/lovemew67/project-misc/grpc-gateway-1/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
