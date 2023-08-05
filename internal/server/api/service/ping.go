// @BasePath /api/v1
package service

import (
	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func PingView(c *gin.Context) {
	c.JSON(200, 
		gin.H{
			"message": "pong",
		})
}