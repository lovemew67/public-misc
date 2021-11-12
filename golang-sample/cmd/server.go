package cmd

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	quit = make(chan os.Signal, 5)
)

func NewAPIServerCmd() *cobra.Command {
	var (
		serverConfigFile string
	)

	var apiServerCmd = &cobra.Command{
		Use:   "server",
		Short: "server",
		Run: func(cmd *cobra.Command, args []string) {
			// shut sonar
			_ = cmd.Runnable()
			_ = len(args)

			// init viper
			viper.AutomaticEnv()
			viper.SetConfigFile(serverConfigFile)
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
			if errViper := viper.ReadInConfig(); errViper != nil {
				log.Printf("[%s] viper read file error: %+v", cornerstone.GetAppName(), errViper)
			}
			log.SetFlags(log.LstdFlags)

			// log config before initializing
			jsIndent, _ := json.MarshalIndent(viper.AllSettings(), "", "\t")
			log.Printf("[%s] initializing server with config: %s", cornerstone.GetAppName(), (jsIndent))

			// add graceful shutdown
			signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

			// blocking
			sig := <-quit
			cornerstone.Infof(systemCtx, "[%s] receive exit signal: %+v", cornerstone.GetAppName(), sig)
		},
	}

	apiServerCmd.Flags().StringVarP(&serverConfigFile, "config", "c", "./local.toml", "Path to Config File")
	return apiServerCmd
}
