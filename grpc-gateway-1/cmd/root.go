package cmd

import (
	"log"

	"github.com/lovemew67/cornerstone"
	"github.com/spf13/cobra"
)

var (
	systemCtx cornerstone.Context
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Llongfile)
	systemCtx = cornerstone.NewContext()
}

var rootCmd = &cobra.Command{
	Use:   "clean-architecture-demo",
	Short: "clean-architecture-demo",
	Run: func(cmd *cobra.Command, args []string) {
		// shut sonar
		_ = cmd.Runnable()
		_ = len(args)
	},
}

func Execute() error {
	rootCmd.AddCommand(NewAPIServerCmd())
	return rootCmd.Execute()
}
