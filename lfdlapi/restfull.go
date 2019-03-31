package lfdlapi

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the router so it can be used in the main and in testing
func setupRouter() *gin.Engine {
	router := gin.Default()

	authMiddleware, _ := authMiddleware()

	router.POST("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), NoRouteHandler)

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

// login data that binds to the data submitted
type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"
