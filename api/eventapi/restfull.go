package eventapi

import (
	"github.com/gin-gonic/gin"

	"github.com/lookfirstdrivelater/lfdlapi/internal/handlers"
)

// SetupRouter sets up gin engine for the api
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.NoRoute(handlers.NoRouteHandler)

	router.GET("/healthcheck", handlers.HealthcheckHandler)

	return router
}
