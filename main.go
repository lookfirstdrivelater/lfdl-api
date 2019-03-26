package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/rollbar"
	"github.com/joho/godotenv"
	"github.com/stvp/roll"
)

func init() {

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := setupRouter()

	roll.Token = os.Getenv("RollBarToken")
	roll.Endpoint = os.Getenv("RollBarEnv")
	//roll.Environment = "production" // defaults to "development"
	router.Use(rollbar.Recovery(true))
	roll.Info("Starting server", map[string]string{})
	fmt.Println("Useing Rollbar reporting")

	// staring the server
	router.Run(":8080")
}
