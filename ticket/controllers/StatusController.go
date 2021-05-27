package controllers

import (
	"net/http"
	"ticket/model"
	Action "ticket/model-action"

	"github.com/gin-gonic/gin"
)

//GetStatusAll - Get all Status
func GetStatusAll(c *gin.Context) {
	var status []model.Status
	err := Action.GetAllStatus(&status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "ok", "data": status})
	}

}

//GetTicketByStatusID - Get the Ticket  by StatusID
func GetTicketByStatusID(c *gin.Context) {
	id := c.Params.ByName("id")
	var ticket []model.Ticket
	err := Action.GetTicketByStatusID(&ticket, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "ok", "data": ticket})
	}
}
