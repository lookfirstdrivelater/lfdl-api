package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowedMethods = []string{"GET", "POST"}

	router.Use(cors.New(config))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(autotls.Run(router, "stupidcpu.com"))

}
