package controllers

import (
	"net/http"
	"ticket/interceptor"
	"ticket/model"
	Action "ticket/model-action"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login - User
func Login(c *gin.Context) {
	var user model.User

	if c.ShouldBind(&user) == nil {
		var queryUser model.User
		if err := Action.CheckHasUser(&queryUser, user.Username); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": "nok", "error": err})
		} else if checkPasswordHash(user.Password, queryUser.Password) == false {
			c.JSON(http.StatusOK, gin.H{"result": "nok", "error": "invalid password"})
		} else {
			token := interceptor.JwtSign(queryUser)

			c.JSON(http.StatusUnauthorized, gin.H{"result": "ok", "token": token})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{"status": "unable to bind data"})
	}
}

// Register - User
func Register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		// user.CreatedAt = time.Now()
		if err := Action.CreateUser(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": "nok", "error": err})
		} else {
			c.JSON(http.StatusCreated, gin.H{"result": "ok", "data": user})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unable to bind data"})
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
