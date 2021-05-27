package api

import (
	"ticket/controllers"

	"github.com/gin-gonic/gin"
)

// SetupStatusAPI - Status of ticket
func SetupStatusAPI(router *gin.Engine) {
	contactAPI := router.Group("/api/")
	{
		// http://localhost/api/status , Get status list
		contactAPI.GET("/status", controllers.GetStatusAll)

		// http://localhost/api/status/XXX , Get ticket by StatusID
		contactAPI.GET("status/:id", controllers.GetTicketByStatusID)
	}
}
