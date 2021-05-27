package api

import (
	"ticket/controllers"

	"github.com/gin-gonic/gin"
)

// SetupUserAPI - Users [CRUD]
func SetupUserAPI(router *gin.Engine) {
	userAPI := router.Group("/api/")
	{
		// http://localhost/api/user , Get users list
		userAPI.GET("/user", controllers.GetUsers)

		// http://localhost/api/user , Create user
		userAPI.POST("user", controllers.CreateUser)

		// http://localhost/api/user/XXX , Get user by ID
		userAPI.GET("user/:id", controllers.GetUserByID)

		// http://localhost/api/user/XXX , Update user by ID
		userAPI.PUT("user/:id", controllers.UpdateUser)

		// http://localhost/api/user/XXX , Delete user by ID
		userAPI.DELETE("user/:id", controllers.DeleteUser)

	}
}
