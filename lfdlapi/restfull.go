package lfdlapi

import (
	"github.com/gin-gonic/gin"
)



// SetupRouter sets up the router so it can be used in the main and in testing
func setupRouter() *gin.Engine {
	router := gin.Default()

	authMiddleware, _ := authMiddleware(db)

	router.POST("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), noRouteHandler)

	authRouter := router.Group("/auth")
	// Refresh time can be longer than token timeout
	authRouter.GET("/refresh_token", authMiddleware.RefreshHandler)
	authRouter.Use(authMiddleware.MiddlewareFunc())
	{ // all our auth routes
		// authRouter.GET("/hello", helloHandler)
	}

	// not auth routes
	router.GET("/ping", pingHandler)

	return router
}

