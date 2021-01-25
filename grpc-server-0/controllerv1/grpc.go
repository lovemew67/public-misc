package controllerv1

import (
	"context"
	"fmt"
	"net"
	"runtime/debug"
	"sync"
	"time"

	gph "github.com/kazegusuri/grpc-panic-handler"

	"github.com/lovemew67/cornerstone"
	"github.com/lovemew67/project-misc/grpc-server-0/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type GrpcServer struct {
	ctx    cornerstone.Context
	port   string
	lock   sync.Mutex
	server *grpc.Server

	proto.UnimplementedEchoServiceServer
}

func (s *GrpcServer) Serve() (err error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return
	}

	gph.InstallPanicHandler(func(r interface{}) {
		cornerstone.Errorf(s.ctx, "[grpc panic recover] error:%v\n %s", err, debug.Stack())
	})

	keepaliveEnforcementPolicy := grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	})
	keepaliveParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:    5 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout: 1 * time.Second, // Wait 1 second for the ping ack before assuming the connection is dead
	})

	uIntOpt := grpc.UnaryInterceptor(gph.UnaryPanicHandler)
	sIntOpt := grpc.StreamInterceptor(gph.StreamPanicHandler)

	gs := grpc.NewServer(uIntOpt, sIntOpt, keepaliveParams, keepaliveEnforcementPolicy)
	proto.RegisterEchoServiceServer(gs, s)

	s.lock.Lock()
	s.server = gs
	s.lock.Unlock()

	err = gs.Serve(lis)
	return
}

func (s *GrpcServer) Close() {
	s.lock.Lock()
	if s.server != nil {
		s.server.GracefulStop()
	}
	s.lock.Unlock()
	cornerstone.Infof(s.ctx, "[grpc server] graceful closed")
}

func (s *GrpcServer) Echo(ctx context.Context, input *proto.HiRequest) (result *proto.HiResponse, err error) {
	result = &proto.HiResponse{
		Success: true,
		Message: fmt.Sprintf("echo: %s", input.Message),
	}
	return
}

func InitGrpcServer(ctx cornerstone.Context, port string) (gs *GrpcServer) {
	gsCtx := ctx.CopyContext()
	gsCtx.Set("worker", "grpc")
	gs = &GrpcServer{
		ctx:  gsCtx,
		port: port,
	}
	return
}
