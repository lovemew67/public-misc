package controllerv1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
	"github.com/lovemew67/public-misc/golang-sample/repositoryv1"
	"github.com/lovemew67/public-misc/golang-sample/servicev1"
	"github.com/spf13/viper"
)

const (
	pathID = "id"
)

var (
	ctx = cornerstone.NewContext()
)

func InitGinServer(_ss servicev1.StaffV1Service, _jr repositoryv1.JobV1Repository, _sr repositoryv1.ScheduleV1Repository) (gs *GinServer) {
	// create gin  server.
	gin.SetMode(viper.GetString("rest.mode"))
	gs = &GinServer{
		ss: _ss,

		jr: _jr,
		sr: _sr,

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
	ss servicev1.StaffV1Service

	jr repositoryv1.JobV1Repository
	sr repositoryv1.ScheduleV1Repository

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

	// add staff v1 handlers
	staffGroup := rootGroup.Group("/v1/staff")
	{
		staffGroup.GET("", gs.listStaffV1Handler)
		staffGroup.POST("", gs.createStaffV1Handler)
		staffGroup.GET(fmt.Sprintf("/:%s", pathID), gs.getStaffV1Handler)
		staffGroup.PATCH(fmt.Sprintf("/:%s", pathID), gs.patchStaffV1Handler)
		staffGroup.DELETE(fmt.Sprintf("/:%s", pathID), gs.deleteStaffV1Handler)
	}

	// add schedule v1 handlers
	scheduleGroup := rootGroup.Group("/v1/schedule")
	{
		scheduleGroup.GET("", gs.listScheduleV1Handler)
		scheduleGroup.POST("", gs.createScheduleV1Handler)
		scheduleGroup.GET(fmt.Sprintf("/:%s", pathID), gs.getScheduleV1Handler)
		scheduleGroup.PATCH(fmt.Sprintf("/:%s", pathID), gs.patchScheduleV1Handler)
		scheduleGroup.DELETE(fmt.Sprintf("/:%s", pathID), gs.deleteScheduleV1Handler)
	}

	// add job v1 handlers
	jobGroup := rootGroup.Group("/v1/job")
	{
		jobGroup.GET("", gs.listJobV1Handler)
		jobGroup.POST("", gs.createJobV1Handler)
		jobGroup.GET(fmt.Sprintf("/:%s", pathID), gs.getJobV1Handler)
		jobGroup.PATCH(fmt.Sprintf("/:%s", pathID), gs.patchJobV1Handler)
		jobGroup.DELETE(fmt.Sprintf("/:%s", pathID), gs.deleteJobV1Handler)
	}
}

func (gs *GinServer) version(c *gin.Context) {
	c.JSON(http.StatusOK, cornerstone.GetVersion())
}

func (gs *GinServer) config(c *gin.Context) {
	c.JSON(http.StatusOK, viper.AllSettings())
}

// staff v1 handlers

func (gs *GinServer) createStaffV1Handler(c *gin.Context) {
	input := &domainv1.CreateStaffV1ServiceRequest{}
	if errBind := c.ShouldBindJSON(input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	result, err := gs.ss.CreateStaffV1Service(input)
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
	result, err := gs.ss.GetStaffV1Service(input)
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
	results, total, err := gs.ss.ListStaffV1Service(input)
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
	err := gs.ss.PatchStaffV1Service(input)
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
	err := gs.ss.DeleteStaffV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, nil)
}

// job v1 handlers

func (gs *GinServer) createJobV1Handler(c *gin.Context) {
	input := &domainv1.CreateJobV1Request{}
	if errBind := c.ShouldBindJSON(input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	result, err := gs.jr.CreateJob(input.Job)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, result)
}

func (gs *GinServer) getJobV1Handler(c *gin.Context) {
	jobID := c.Param(pathID)
	result, err := gs.jr.GetJob(jobID)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, result)
}

func (gs *GinServer) listJobV1Handler(c *gin.Context) {
	input := &domainv1.ListJobV1Request{}
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
	total, err := gs.jr.CountTotalJobs()
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	results, err := gs.jr.QueryAllJobsWithOffsetAndLimit(input.Offset, input.Limit)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, gin.H{
		"staff": results,
		"total": total,
	})
}

func (gs *GinServer) patchJobV1Handler(c *gin.Context) {
	input := &domainv1.PatchJobV1Request{}
	if errBind := c.ShouldBindJSON(input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	jobID := c.Param(pathID)
	updater := &domainv1.Job{}
	if input.RetryCount != nil {
		updater.RetryCount = *input.RetryCount
	}
	if input.Status != nil {
		updater.Status = *input.Status
	}
	if input.Type != nil {
		updater.Type = *input.Type
	}
	if input.Processing != nil {
		updater.Processing = *input.Processing
	}
	err := gs.jr.PatchJob(jobID, updater)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, nil)
}

func (gs *GinServer) deleteJobV1Handler(c *gin.Context) {
	jobID := c.Param(pathID)
	err := gs.jr.DeleteJob(jobID)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, nil)
}

// schedule v1 handlers

func (gs *GinServer) createScheduleV1Handler(c *gin.Context) {
	input := &domainv1.CreateScheduleV1Request{}
	if errBind := c.ShouldBindJSON(input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	result, err := gs.sr.CreateSchedule(input.Schedule)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, result)
}

func (gs *GinServer) getScheduleV1Handler(c *gin.Context) {
	scheduleID := c.Param(pathID)
	result, err := gs.sr.GetSchedule(scheduleID)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, result)
}

func (gs *GinServer) listScheduleV1Handler(c *gin.Context) {
	input := &domainv1.ListScheduleV1Request{}
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
	total, err := gs.sr.CountTotalSchedules()
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	results, err := gs.sr.QueryAllSchedulesWithOffsetAndLimit(input.Offset, input.Limit)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, gin.H{
		"staff": results,
		"total": total,
	})
}

func (gs *GinServer) patchScheduleV1Handler(c *gin.Context) {
	input := &domainv1.PatchScheduleV1Request{}
	if errBind := c.ShouldBindJSON(input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	scheduleID := c.Param(pathID)
	updater := &domainv1.Schedule{}
	if input.TimeInHourUTC != nil {
		updater.TimeInHourUTC = *input.TimeInHourUTC
	}
	if input.Enable != nil {
		updater.Enable = *input.Enable
	}
	if input.Type != nil {
		updater.Type = *input.Type
	}
	err := gs.sr.PatchSchedule(scheduleID, updater)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, nil)
}

func (gs *GinServer) deleteScheduleV1Handler(c *gin.Context) {
	scheduleID := c.Param(pathID)
	err := gs.sr.DeleteSchedule(scheduleID)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, nil)
}
