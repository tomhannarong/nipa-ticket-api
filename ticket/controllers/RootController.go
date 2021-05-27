package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetRoot - Welcome
func GetRoot(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"result": "Welcome Ticket Management Application."})
}
