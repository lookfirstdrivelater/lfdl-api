package main

import (
	"fmt"
	"log"
	"os"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/rollbar"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stvp/roll"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()

	roll.Token = os.Getenv("RollBarToken")
	roll.Endpoint = os.Getenv("RollBarEnv")
	//roll.Environment = "production" // defaults to "development"
	router.Use(rollbar.Recovery(true))
	roll.Info("Starting server", map[string]string{})
	fmt.Println("Useing Rollbar reporting")

	authMiddleware, _ := authMiddleware()

	router.POST("/login", authMiddleware.LoginHandler)
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

	// staring the server
	router.Run(":8080")
}
