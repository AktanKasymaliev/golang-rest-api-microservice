package api

import (
	"base/internal/server/api/endpoints"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouteRegister(router *gin.Engine) {
	router.GET("/swagger/", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/api/v1/ping", endpoints.PingView)
	router.GET("/api/v1/users", endpoints.UsersView)
}
