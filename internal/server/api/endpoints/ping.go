package endpoints

import (
	_ "base/docs"  //

	"net/http"

	"github.com/gin-gonic/gin"
)

// PingView 			godoc
// @Summary 			Show the status of server.
// @Description 		Get the status of server.
// @Tags 				Ping
// @Accept 				*/*
// @Produce 			json
// @Success 			200 {object} map[string]interface{}
// @Router 				/ping [get]
func PingView(c *gin.Context) {
	c.JSON(http.StatusOK, 
		gin.H{
			"message": "pong",
		})
}