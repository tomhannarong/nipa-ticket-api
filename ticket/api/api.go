package api

import (
	"ticket/db"

	"github.com/gin-gonic/gin"
)

// Setup - API
func Setup(router *gin.Engine) {

	db.SetupDB()
	SetupRootAPI(router)
	SetupAuthenAPI(router)
	SetupUserAPI(router)
	SetupTicketAPI(router)
	SetupStatusAPI(router)

}
