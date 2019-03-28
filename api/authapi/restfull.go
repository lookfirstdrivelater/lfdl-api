package authapi

import (
	"github.com/lookfirstdrivelater/lfdlapi/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/lookfirstdrivelater/lfdlapi/internal/middleware"
)

// SetupRouter sets up the router so it can be used in the main and in testing
func SetupRouter() *gin.Engine {
	router := gin.Default()

	authMiddleware, _ := middleware.AuthMiddleware()

	router.POST("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), handlers.NoRouteHandler)

	authRouter := router.Group("/auth")
	// Refresh time can be longer than token timeout
	authRouter.GET("/refresh_token", authMiddleware.RefreshHandler)
	authRouter.Use(authMiddleware.MiddlewareFunc())
	{ // all our auth routes
		authRouter.GET("/hello", helloHandler)
	}

	// not auth routes
	router.GET("/ping", pingHandler)

	return router
}
