package authapi

import (
	"time"

	internalauthapi "github.com/lookfirstdrivelater/lfdlapi/internal/authapi"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"userName": user.(*internalauthapi.User).UserName,
		"text":     "Hello World.",
	})
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":       "pong",
		"systemTimeUTC": time.Now(),
	})
}
