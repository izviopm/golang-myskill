package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const(
	USER="admin"
	PASSWORD="Password123!"
	SECRET="secret"
)

func AuthValid(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Token Required",
		})
		c.Abort()
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, invalid := token.Method.(*jwt.SigningMethodHMAC); !invalid{
			return nil, fmt.Errorf("Invalid token ", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})

	if token != nil && err == nil {
		fmt.Println("Token Verified")
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
			"error": err.Error(),
		})
		c.Abort()
	}
}