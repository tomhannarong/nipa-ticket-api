package main

import (
	"fmt"
	"os"
	"ticket/api"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	// Set Gin
	router := gin.Default()

	// Set helmet
	router.Use(helmet.Default())
	router.Use(helmet.NoCache())

	// Set up CORS middleware options
	config := cors.Config{
		Origins:         "*",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		Methods:         "GET, POST, PUT, DELETE",
		Credentials:     true,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}

	// Apply the middleware to the router (works on groups too)
	router.Use(cors.Middleware(config))

	api.Setup(router)

	// Start Set Port
	// set port on docker OR production
	var port = os.Getenv("PORT")
	if port == "" {
		// in case on Local port default
		fmt.Println("No Port In Production") // :8080
		router.Run()
	} else {
		// in case on production
		fmt.Println("Environment Port : " + port)
		router.Run(fmt.Sprintf(":%s", port))
	}
}
