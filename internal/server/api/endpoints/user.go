package endpoints

import (
	"base/internal/server/api/service"

	"github.com/gin-gonic/gin"
)


func UsersView(c *gin.Context) {
	service.UserResponse(c)
}