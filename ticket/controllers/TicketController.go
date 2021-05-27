package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"ticket/model"
	Action "ticket/model-action"
	"time"

	"github.com/gin-gonic/gin"
)

//GetTickets - Get all Ticket
func GetTickets(c *gin.Context) {
	var ticket []model.Ticket
	queryString := c.Query("status")
	// fmt.Println(queryString)
	if queryString == "" {
		err := Action.GetAllTickets(&ticket)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "ok", "data": ticket})
		}
	} else {
		err := Action.GetFilterTickets(&ticket, queryString)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "ok", "data": ticket})
		}
	}

}

//CreateTicket - Create Ticket
func CreateTicket(c *gin.Context) {
	var ticket model.Ticket
	// var contact model.Contact
	// Step - Create Ticket
	if c.ShouldBind(&ticket) == nil {
		ticket.CreatedAt = time.Now()
		err := Action.CreateTicket(&ticket)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"result": "nok", "error": err})
		} else {

			Action.GetTicketByID(&ticket, strconv.FormatUint(ticket.ID, 10))
			c.JSON(http.StatusCreated, gin.H{"result": "ok", "data": ticket})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "unable to bind data"})
	}
}

//GetTicketByID - Get the Ticket by id
func GetTicketByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var ticket model.Ticket
	err := Action.GetTicketByID(&ticket, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "nok", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "ok", "data": ticket})
	}
}

//UpdateTicket - Update the Ticket information
func UpdateTicket(c *gin.Context) {
	var ticket model.Ticket

	id := c.Params.ByName("id")
	// Get Ticket By ID
	err := Action.GetTicketByID(&ticket, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
	}
	// Update Ticket
	if c.ShouldBind(&ticket) == nil {
		StatusID, _ := strconv.ParseUint(c.PostForm("StatusID"), 10, 64)
		ticket.Status.ID = StatusID
		fmt.Println(ticket)
		err := Action.UpdateTicket(&ticket, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
		} else {
				// Action.GetTicketByID(&ticket, id)
				c.JSON(http.StatusOK, gin.H{"result": "ok", "data": ticket})		
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "unable to bind data"})
	}
}

//DeleteTicket - Delete the Ticket
func DeleteTicket(c *gin.Context) {
	var ticket model.Ticket
	id := c.Params.ByName("id")
	err := Action.DeleteTicket(&ticket, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
