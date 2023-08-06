package endpoints

import (
	"github.com/gin-gonic/gin"
)

func PingView(c *gin.Context) {
	c.JSON(200, 
		gin.H{
			"message": "pong",
		})
}