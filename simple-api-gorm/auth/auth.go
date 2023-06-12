package auth

import (
	"net/http"
	"simple-api-gorm/models"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

const(
	USER="admin"
	PASSWORD="Password123!"
	SECRET="secret"
)

func LoginHandler(c *gin.Context) {
	var user models.Credential

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	if user.Username != USER {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User Invalid",
		})
	} else if user.Password != PASSWORD {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password Invalid",
		})
	} else {
		claim := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute*1).Unix(),
			Issuer: "test",
			IssuedAt: time.Now().Unix(),
		}
	
		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		token, err := sign.SignedString([]byte(SECRET))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
		}
	
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"token": token,
		})
	}

}