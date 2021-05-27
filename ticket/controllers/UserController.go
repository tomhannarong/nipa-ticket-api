package controllers

import (
	"fmt"
	"net/http"
	"ticket/model"
	Action "ticket/model-action"

	"github.com/gin-gonic/gin"
)

//GetUsers - Get all users
func GetUsers(c *gin.Context) {
	var user []model.User
	err := Action.GetAllUsers(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "ok", "data": user})
	}
}

//CreateUser - Create User
func CreateUser(c *gin.Context) {
	var user model.User
	// c.BindJSON(&user)
	if c.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		err := Action.CreateUser(&user)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"result": "nok", "error": err})
		} else {
			c.JSON(http.StatusCreated, gin.H{"result": "ok", "data": user})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unable to bind data"})
	}
}

//GetUserByID - Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user model.User
	err := Action.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "ok", "data": user})
	}
}

//UpdateUser - Update the user information
func UpdateUser(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := Action.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
	}
	if c.ShouldBind(&user) == nil {
		err = Action.UpdateUser(&user, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "ok", "data": user})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unable to bind data"})
	}
}

//DeleteUser - Delete the user
func DeleteUser(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	err := Action.DeleteUser(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "nok", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
