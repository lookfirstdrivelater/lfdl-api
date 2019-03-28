//Package handlers ...
package handlers

import (
	"log"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// HealthcheckHandler is a general healthcheck handler
func HealthcheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"ServiceName": "xyz",
		"Version":     "0.0.1",
		"health":      "good",
	})
}

// NoRouteHandler is a general no route handler
func NoRouteHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
