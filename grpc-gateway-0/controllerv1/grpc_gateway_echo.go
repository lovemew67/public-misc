package controllerv1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lovemew67/cornerstone"
	"github.com/lovemew67/project-misc/grpc-gateway-0/gen/proto"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

func InitEchoGrpcGateway(ctx cornerstone.Context, grpcPort, grpcGateway string) (gwServer *http.Server) {
	funcName := "InitEchoGrpcGateway"

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
	errRegister := proto.RegisterEchoServiceHandler(context.Background(), gwMux, grpcConn)
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
