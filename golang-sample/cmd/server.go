package cmd

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/controllerv1"
	"github.com/lovemew67/public-misc/golang-sample/repositoryv1/sqlite"
	"github.com/lovemew67/public-misc/golang-sample/servicev1"
	"github.com/lovemew67/public-misc/golang-sample/workerv1"
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
			funcName := "apiServerCmd.Run"

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

			// init repository
			staffV1Repositorier, errRepository := sqlite.NewStaffV1SQLiteRepositorier(systemCtx)
			if errRepository != nil {
				cornerstone.Panicf(systemCtx, "[%s] failed to create staff v1 repositiory, err: %+v", funcName, errRepository)
			}

			// init service
			staffV1Service, errService := servicev1.NewStaffV1Servicer(staffV1Repositorier)
			if errService != nil {
				cornerstone.Panicf(systemCtx, "[%s] failed to create staff v1 service, err: %+v", funcName, errService)
			}

			// init grpc server
			grpcServer := controllerv1.InitGRPCServer(staffV1Service)
			grpcCanceller := controllerv1.GRPCListenAndServe(systemCtx, grpcServer)
			defer grpcCanceller()

			// init http server
			ginServer := controllerv1.InitGinServer(staffV1Service)
			ginCanceller := controllerv1.HTTPListenAndServe(systemCtx, ginServer)
			defer ginCanceller()

			// init grpc gateway
			grpcGateway := controllerv1.InitGRPCGateway(systemCtx)
			grpcGatewayCanceller := controllerv1.GRPCGatewayListenAndServe(systemCtx, grpcGateway)
			defer grpcGatewayCanceller()

			// init scheduler
			scheduleTicker := workerv1.InitScheduler(systemCtx)
			defer scheduleTicker.Stop()

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
