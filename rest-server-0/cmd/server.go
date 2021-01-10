package cmd

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/lovemew67/cornerstone"
	"github.com/lovemew67/project-misc/rest-server-0/controllerv1"
	"github.com/lovemew67/project-misc/rest-server-0/repositoryv1"
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
				log.Panicf("[%s] viper read file error: %+v", cornerstone.GetAppName(), errViper)
			}

			// log config before initializing
			jsIndent, _ := json.MarshalIndent(viper.AllSettings(), "", "\t")
			log.Printf("[%s] initializing server with config: %s", cornerstone.GetAppName(), (jsIndent))

			// init repository
			repositoryv1.Init(systemCtx)

			// init http server
			port := viper.GetString("http.port")
			router := controllerv1.InitGinServer()
			httpServer := &http.Server{
				Addr:         ":" + port,
				Handler:      router,
				ReadTimeout:  viper.GetDuration("http.read_timeout"),
				WriteTimeout: viper.GetDuration("http.write_timeout"),
			}
			go func() {
				cornerstone.Infof(systemCtx, "[%s] server is running and listening port: %s", cornerstone.GetAppName(), port)
				if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					cornerstone.Panicf(systemCtx, "[%s] failed to listen on port: %s, err: %+v", cornerstone.GetAppName(), port, err)
				}
			}()
			defer func() {
				cornerstone.Infof(systemCtx, "[%s] shuting down server", cornerstone.GetAppName())
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if errShutdown := httpServer.Shutdown(ctx); errShutdown != nil {
					cornerstone.Panicf(systemCtx, "[%s] failed to shut down server, err: %+v", cornerstone.GetAppName(), errShutdown)
				}
				cornerstone.Infof(systemCtx, "[%s] server exiting", cornerstone.GetAppName())
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
