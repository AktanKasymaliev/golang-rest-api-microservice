package api

import (
	"base/internal/server/api/endpoints"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// RouteRegister regiter all endpoints
func RouteRegister(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/ping/", endpoints.PingView)

	// Users
	users := v1.Group("users/")
	users.GET("/get-all/", endpoints.UsersView)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
