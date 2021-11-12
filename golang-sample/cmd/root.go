package cmd

import (
	"log"

	"github.com/lovemew67/public-misc/cornerstone"
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
	Use:   "golang-sample",
	Short: "golang-sample",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	rootCmd.AddCommand(NewAPIServerCmd())
	return rootCmd.Execute()
}
