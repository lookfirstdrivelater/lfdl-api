package eventapi

import (
	"log"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up gin engine for the api
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.GET("/healthcheck", healthcheck)

	return router
}
