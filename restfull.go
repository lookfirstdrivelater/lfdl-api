package main

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

type users struct {
	gorm.Model
	UserName  string
	FirstName string
	LastName  string
	Password  string
}

type event struct {
	gorm.Model
	StartTime time.Time
	EndTime   time.Time
}

// SetupRouter sets up the router so it can be used in the main and in testing
func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	authMiddleware, _ := authMiddleware(db)

	router.POST("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), NoRouteHandler)

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
