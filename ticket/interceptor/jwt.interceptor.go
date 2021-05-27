package interceptor

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"ticket/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JwtSign - Get Token
func JwtSign(payload model.User) string {
	atClaims := jwt.MapClaims{}

	// Payload begin
	atClaims["id"] = payload.ID
	atClaims["username"] = payload.Username
	atClaims["level"] = payload.Level
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// Payload end

	// Set JWT
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return token
}

// JwtVerify - Token Verify
func JwtVerify(c *gin.Context) {
	// Step Get Token from header request
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	fmt.Println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	// Step Verify Token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check Token Pass Verify
		fmt.Println(claims)

		staffID := fmt.Sprintf("%v", claims["id"])
		username := fmt.Sprintf("%v", claims["jwt_username"]) //get value JSON from JWT token
		level := fmt.Sprintf("%v", claims["jwt_level"])       //get value JSON from JWT token
		c.Set("jwt_staff_id", staffID)                        // set value jwt to context
		c.Set("jwt_username", username)                       // set value jwt to context
		c.Set("jwt_level", level)                             // set value jwt to context

		c.Next()
	} else {
		// Check Token Not Pass Verify
		c.JSON(http.StatusOK, gin.H{"result": "nok", "message": "invalid token", "error": err})
		c.Abort()
	}
}
