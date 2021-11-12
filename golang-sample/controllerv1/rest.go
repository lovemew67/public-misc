package controllerv1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/grpc-gateway-1/domainv1"
	"github.com/lovemew67/public-misc/grpc-gateway-1/servicev1"
	"github.com/spf13/viper"
)

const (
	pathID = "id"
)

func InitGinServer(_s servicev1.StaffV1Service) (gs *GinServer) {
	// create gin  server.
	gin.SetMode(viper.GetString("rest.mode"))
	gs = &GinServer{
		s:      _s,
		Engine: gin.New(),
	}
	gs.initRoutings()
	return
}

func HTTPListenAndServe(ctx cornerstone.Context, gs *GinServer) (canceller func()) {
	funcName := "HTTPListenAndServe"
	restPort := viper.GetString("rest.port")
	httpServer := &http.Server{
		Addr:         ":" + restPort,
		Handler:      gs,
		ReadTimeout:  viper.GetDuration("rest.read_timeout"),
		WriteTimeout: viper.GetDuration("rest.write_timeout"),
	}
	go func() {
		cornerstone.Infof(ctx, "[%s] http server is running and listening port: %s", funcName, restPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			cornerstone.Panicf(ctx, "[%s] http server failed to listen on port: %s, err: %+v", funcName, restPort, err)
		}
	}()

	routineCtx := ctx.CopyContext()
	canceller = func() {
		cornerstone.Infof(routineCtx, "[%s] shuting down http server", cornerstone.GetAppName())
		nativeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if errShutdown := httpServer.Shutdown(nativeCtx); errShutdown != nil {
			cornerstone.Panicf(routineCtx, "[%s] failed to shut down http server, err: %+v", cornerstone.GetAppName(), errShutdown)
		}
		cornerstone.Infof(routineCtx, "[%s] http server exiting", cornerstone.GetAppName())
	}
	return
}

type GinServer struct {
	s servicev1.StaffV1Service

	*gin.Engine
}

func (gs *GinServer) initRoutings() {
	// add data retention group
	rootGroup := gs.Group("")

	// general service for debugging
	{
		rootGroup.GET("/config", gs.config)
		rootGroup.GET("/version", gs.version)
	}

	// add v1 handlers
	staffGroup := rootGroup.Group("/v1/staff")
	{
		staffGroup.GET("", gs.listStaffV1Handler)
		staffGroup.POST("", gs.createStaffV1Handler)
		staffGroup.GET(fmt.Sprintf("/:%s", pathID), gs.getStaffV1Handler)
		staffGroup.PATCH(fmt.Sprintf("/:%s", pathID), gs.patchStaffV1Handler)
		staffGroup.DELETE(fmt.Sprintf("/:%s", pathID), gs.deleteStaffV1Handler)
	}
}

func (gs *GinServer) version(c *gin.Context) {
	c.JSON(http.StatusOK, cornerstone.GetVersion())
}

func (gs *GinServer) config(c *gin.Context) {
	c.JSON(http.StatusOK, viper.AllSettings())
}

func (gs *GinServer) createStaffV1Handler(c *gin.Context) {
	input := &domainv1.CreateStaffV1ServiceRequest{}
	if errBind := c.ShouldBindJSON(input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	result, err := gs.s.CreateStaffV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, result)
}

func (gs *GinServer) getStaffV1Handler(c *gin.Context) {
	staffID := c.Param(pathID)
	input := &domainv1.GetStaffV1ServiceRequest{
		ID: staffID,
	}
	result, err := gs.s.GetStaffV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, result)
}

func (gs *GinServer) listStaffV1Handler(c *gin.Context) {
	input := &domainv1.ListStaffV1ServiceRequest{}
	if errBind := c.BindQuery(&input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	if input.Limit <= 0 {
		input.Limit = 10
	}
	if input.Limit > 200 {
		input.Limit = 200
	}
	results, total, err := gs.s.ListStaffV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, gin.H{
		"staff": results,
		"total": total,
	})
}

func (gs *GinServer) patchStaffV1Handler(c *gin.Context) {
	input := &domainv1.PatchStaffV1ServiceRequest{}
	if errBind := c.ShouldBindJSON(input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	staffID := c.Param(pathID)
	input.ID = staffID
	err := gs.s.PatchStaffV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, nil)
}

func (gs *GinServer) deleteStaffV1Handler(c *gin.Context) {
	staffID := c.Param(pathID)
	input := &domainv1.DeleteStaffV1ServiceRequest{
		ID: staffID,
	}
	err := gs.s.DeleteStaffV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, nil)
}
