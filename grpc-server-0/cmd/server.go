package cmd

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/lovemew67/cornerstone"
	"github.com/lovemew67/project-misc/grpc-server-0/controllerv1"
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

			// init grpc server
			grpcPort := viper.GetString("grpc.port")
			grpcServer := controllerv1.InitGrpcServer(systemCtx, grpcPort)
			defer func() {
				cornerstone.Infof(systemCtx, "[%s] shuting down grpc server", cornerstone.GetAppName())
				grpcServer.Close()
			}()
			go func() {
				cornerstone.Infof(systemCtx, "[%s] grpc server is running and listening port: %s", cornerstone.GetAppName(), grpcPort)
				if errServe := grpcServer.Serve(); errServe != nil {
					cornerstone.Panicf(systemCtx, "[%s] grpc serve error: %s", cornerstone.GetAppName(), errServe)
				}
			}()

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
