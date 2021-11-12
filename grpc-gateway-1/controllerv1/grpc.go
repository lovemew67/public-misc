package controllerv1

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/lovemew67/project-misc/grpc-gateway-1/domainv1"
	"github.com/lovemew67/project-misc/grpc-gateway-1/gen/proto"
	"github.com/lovemew67/project-misc/grpc-gateway-1/servicev1"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func InitGRPCServer(_s servicev1.StaffV1Service) (gs *grpc.Server) {
	// create grpc server
	gs = grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)

	// register grpc service
	proto.RegisterStaffServiceV1Server(gs, &GRPCServer{
		s: _s,
	})
	reflection.Register(gs)

	return
}

func GRPCListenAndServe(ctx cornerstone.Context, gs *grpc.Server) (canceller func()) {
	funcName := "GRPCListenAndServe"
	grpcPort := viper.GetString("grpc.port")
	grpcListener, errListen := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if errListen != nil {
		cornerstone.Panicf(ctx, "[%s] failed to listen: %+v", funcName, errListen)
	}
	go func() {
		cornerstone.Infof(ctx, "[%s] grpc server is running and listening port: %s", funcName, grpcPort)
		if err := gs.Serve(grpcListener); err != nil {
			cornerstone.Panicf(ctx, "[%s] grpc server failed to listen on port: %s, err: %+v", funcName, grpcPort, err)
		}
	}()

	routineCtx := ctx.CopyContext()
	canceller = func() {
		cornerstone.Infof(routineCtx, "[%s] shuting down grpc server", cornerstone.GetAppName())
		gs.GracefulStop()
		cornerstone.Infof(routineCtx, "[%s] grpc server exiting", cornerstone.GetAppName())
	}
	return
}

type GRPCServer struct {
	s servicev1.StaffV1Service

	proto.UnimplementedStaffServiceV1Server
}

func (gs *GRPCServer) ListStaffV1(ctx context.Context, in *proto.ListStaffV1Request) (*proto.ListStaffV1Response, error) {
	input := &domainv1.ListStaffV1ServiceRequest{
		Offset: int(in.Offset),
		Limit:  int(in.Limit),
	}
	if input.Limit <= 0 {
		input.Limit = 10
	}
	if input.Limit > 200 {
		input.Limit = 200
	}
	out := &proto.ListStaffV1Response{}
	results, total, err := gs.s.ListStaffV1Service(input)
	if err != nil {
		out.ErrorMessage = err.Error()
		return out, nil
	}
	out.Total = int32(total)
	out.Staff = results
	return out, nil
}

func (gs *GRPCServer) CreateStaffV1(ctx context.Context, in *proto.CreateStaffV1Request) (*proto.CreateStaffV1Response, error) {
	input := &domainv1.CreateStaffV1ServiceRequest{}
	input.StaffV1 = in.Staff
	out := &proto.CreateStaffV1Response{}
	result, err := gs.s.CreateStaffV1Service(input)
	if err != nil {
		out.ErrorMessage = err.Error()
		return out, nil
	}
	out.Staff = result
	return out, nil
}

func (gs *GRPCServer) GetStaffV1(ctx context.Context, in *proto.GetStaffV1Request) (*proto.GetStaffV1Response, error) {
	input := &domainv1.GetStaffV1ServiceRequest{}
	input.ID = in.Id
	out := &proto.GetStaffV1Response{}
	result, err := gs.s.GetStaffV1Service(input)
	if err != nil {
		out.ErrorMessage = err.Error()
		return out, nil
	}
	out.Staff = result
	return out, nil
}

func (gs *GRPCServer) PatchStaffV1(ctx context.Context, in *proto.PatchStaffV1Request) (*proto.PatchStaffV1Response, error) {
	input := &domainv1.PatchStaffV1ServiceRequest{}
	input.ID = in.Id
	input.Name = &in.Name
	input.Email = &in.Email
	input.AvatarUrl = &in.AvatarUrl
	out := &proto.PatchStaffV1Response{}
	err := gs.s.PatchStaffV1Service(input)
	if err != nil {
		out.ErrorMessage = err.Error()
		return out, nil
	}
	return out, nil
}

func (gs *GRPCServer) DeleteStaffV1(ctx context.Context, in *proto.DeleteStaffV1Request) (*proto.DeleteStaffV1Response, error) {
	input := &domainv1.DeleteStaffV1ServiceRequest{}
	input.ID = in.Id
	out := &proto.DeleteStaffV1Response{}
	err := gs.s.DeleteStaffV1Service(input)
	if err != nil {
		out.ErrorMessage = err.Error()
		return out, nil
	}
	return out, nil
}
