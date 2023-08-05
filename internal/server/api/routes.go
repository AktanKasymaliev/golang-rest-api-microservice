package api

import (
	"base/internal/server/api/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouteRegister(router *gin.Engine) {
	router.GET("/swagger/", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/api/v1/ping",service.PingView)
}
