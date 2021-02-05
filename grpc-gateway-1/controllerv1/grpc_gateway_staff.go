package controllerv1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lovemew67/cornerstone"
	"github.com/lovemew67/project-misc/grpc-gateway-1/gen/proto"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitGRPCGateway(ctx cornerstone.Context) (gwServer *http.Server) {
	funcName := "InitGRPCGateway"

	grpcPort := viper.GetString("grpc.port")
	grpcGateway := viper.GetString("grpc.gateway_port")

	grpcConn, errDial := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%s", grpcPort),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if errDial != nil {
		cornerstone.Panicf(ctx, "[%s] failed to dial grpc conn, err: %+v", funcName, errDial)
	}

	gwMux := runtime.NewServeMux()
	errRegister := proto.RegisterStaffServiceV1Handler(context.Background(), gwMux, grpcConn)
	if errRegister != nil {
		cornerstone.Panicf(ctx, "[%s] failed to register handler, err: %+v", funcName, errRegister)
	}

	// for demo purpose, must add CORS header for swagger UI
	corsContext := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	gwServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", grpcGateway),
		Handler: corsContext.Handler(gwMux),
	}
	return
}

func GRPCGatewayListenAndServe(ctx cornerstone.Context, gwServer *http.Server) (canceller func()) {
	funcName := "GRPCGatewayListenAndServe"
	grpcGatewayPort := viper.GetString("grpc.gateway_port")
	go func() {
		cornerstone.Infof(ctx, "[%s] grpc gateway is running and listening port: %s", funcName, grpcGatewayPort)
		if errServe := gwServer.ListenAndServe(); errServe != nil {
			cornerstone.Panicf(ctx, "[%s] grpc gateway error: %s", funcName, errServe)
		}
	}()

	routineCtx := ctx.CopyContext()
	canceller = func() {
		cornerstone.Infof(routineCtx, "[%s] shuting down grpc gateway", funcName)
		// grpcGateway.Close()
		cornerstone.Infof(routineCtx, "[%s] grpc gateway exiting", funcName)
	}
	return
}
