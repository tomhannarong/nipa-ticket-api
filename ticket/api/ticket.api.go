package api

import (
	"ticket/controllers"

	"github.com/gin-gonic/gin"
)

// SetupTicketAPI - Users [CRUD]
func SetupTicketAPI(router *gin.Engine) {
	ticketAPI := router.Group("/api/")
	{
		// http://localhost/api/ticket , Get ticket list
		ticketAPI.GET("/ticket", controllers.GetTickets)

		// http://localhost/api/ticket , Create ticket
		ticketAPI.POST("ticket", controllers.CreateTicket)

		// http://localhost/api/ticket/XXX , Get ticket by ID
		ticketAPI.GET("ticket/:id", controllers.GetTicketByID)

		// http://localhost/api/ticket/XXX , Update ticket by ID
		ticketAPI.PUT("ticket/:id", controllers.UpdateTicket)

		// http://localhost/api/ticket/XXX , Delete ticket by ID
		ticketAPI.DELETE("ticket/:id", controllers.DeleteTicket)

	}
}
