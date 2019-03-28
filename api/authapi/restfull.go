package authapi

import (
	"log"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	authMiddleware, _ := authMiddleware()

	router.GET("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

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
