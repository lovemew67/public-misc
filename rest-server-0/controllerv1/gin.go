package controllerv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lovemew67/cornerstone"
	"github.com/spf13/viper"
)

func InitGinServer() (router *gin.Engine) {
	// create gin http server.
	gin.SetMode(viper.GetString("http.mode"))
	router = gin.New()

	// add data retention group
	rootGroup := router.Group("")

	// general service for debugging
	{
		rootGroup.GET("/config", config)
		rootGroup.GET("/version", version)
	}

	// add v1 handlers
	addHTTPV1Endpoint(rootGroup)

	return
}

func version(c *gin.Context) {
	c.JSON(http.StatusOK, cornerstone.GetVersion())
}

func config(c *gin.Context) {
	c.JSON(http.StatusOK, viper.AllSettings())
}
