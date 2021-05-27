package api

import (
	"ticket/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRootAPI - Welcome Root
func SetupRootAPI(router *gin.Engine) {
	rootAPI := router.Group("/")
	{
		// Get Root
		rootAPI.GET("/", controllers.GetRoot)
	}
}
