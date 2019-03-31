package lfdlapi

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// healthCheckHandler is a general healthcheck handler
func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"ServiceName": "xyz",
		"Version":     "0.0.1",
		"health":      "good",
	})
}

// noRouteHandler is a general no route handler
func noRouteHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	// user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		// "userName": user.(*user).UserName,
		"text": "Hello World.",
	})
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":       "pong",
		"systemTimeUTC": time.Now(),
	})
}
