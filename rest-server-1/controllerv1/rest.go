package controllerv1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lovemew67/cornerstone"
	"github.com/lovemew67/project-misc/rest-server-1/servicev1"
)

const (
	pathID = "id"
)

func addHTTPV1Endpoint(rootGroup *gin.RouterGroup) {
	staffGroup := rootGroup.Group("/v1/staff")
	{
		staffGroup.GET("", listStaffV1Handler)
		staffGroup.POST("", createStaffV1Handler)
		staffGroup.GET(fmt.Sprintf("/:%s", pathID), getStaffV1Handler)
		staffGroup.PATCH(fmt.Sprintf("/:%s", pathID), patchStaffV1Handler)
		staffGroup.DELETE(fmt.Sprintf("/:%s", pathID), deleteStaffV1Handler)
	}
}

func createStaffV1Handler(c *gin.Context) {
	input := &servicev1.CreateStaffV1ServiceRequest{}
	if errBind := c.ShouldBindJSON(input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	result, err := servicev1.CreateStaffV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, result)
}

func getStaffV1Handler(c *gin.Context) {
	staffID := c.Param(pathID)
	input := &servicev1.GetStaffV1ServiceRequest{
		ID: staffID,
	}
	result, err := servicev1.GetStaffV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, result)
}

func listStaffV1Handler(c *gin.Context) {
	input := &servicev1.ListStaffV1ServiceRequest{}
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
	results, total, err := servicev1.ListStaffV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, gin.H{
		"staff": results,
		"total": total,
	})
}

func patchStaffV1Handler(c *gin.Context) {
	input := &servicev1.PatchStaffV1ServiceRequest{}
	if errBind := c.ShouldBindJSON(input); errBind != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(errBind))
		return
	}
	staffID := c.Param(pathID)
	input.ID = staffID
	err := servicev1.PatchStaffV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, nil)
}

func deleteStaffV1Handler(c *gin.Context) {
	staffID := c.Param(pathID)
	input := &servicev1.DeleteStaffV1ServiceRequest{
		ID: staffID,
	}
	err := servicev1.DeleteStaffV1Service(input)
	if err != nil {
		cornerstone.FromCodeErrorWithStatus(c, cornerstone.FromNativeError(err))
		return
	}
	cornerstone.DoneWithStatus(c, nil)
}
