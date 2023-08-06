package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AllUsersService handles the HTTP request to retrieve all users.
func AllUsersService(c *gin.Context) {
	c.JSON(http.StatusOK, 
		gin.H{
			"message": "ok",
		})	
}