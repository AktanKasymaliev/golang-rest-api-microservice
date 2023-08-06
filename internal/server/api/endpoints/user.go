package endpoints

import (
	_ "base/docs" //
	"base/internal/server/api/services"

	"github.com/gin-gonic/gin"
)

// UsersView 			godoc
// @Summary 			Show the list of users.
// @Description 		Get the all users.
// @Tags 				Users
// @Accept 				*/*
// @Produce 			json
// @Success 			200 {object} map[string]interface{}
// @Router 				/users/get-all [get]
func UsersView(c *gin.Context) {
	services.AllUsersService(c)
}