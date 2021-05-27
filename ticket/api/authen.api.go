package api

import (
	"ticket/controllers"

	"github.com/gin-gonic/gin"
)

// SetupAuthenAPI - login, register
func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/")
	{
		// Login , Get token
		authenAPI.POST("/login", controllers.Login)

		// Register , set user customer .
		authenAPI.POST("/register", controllers.Register)
	}
}
