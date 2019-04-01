package lfdlapi

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// noRouteHandler is a general no route handler
func noRouteHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}

// healthCheckHandler is a general health check handler
func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"ServiceName": "xyz",
		"Version":     "0.0.1",
		"health":      "good",
	})
}

func whoAmI(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
	})
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":       "pong",
		"systemTimeUTC": time.Now(),
	})
}
