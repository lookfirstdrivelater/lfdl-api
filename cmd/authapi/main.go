package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/rollbar"
	"github.com/joho/godotenv"
	"github.com/stvp/roll"

	"github.com/lookfirstdrivelater/lfdlapi/api/authapi"
)

const versionNumber = "0.3.0"

var (
	help       bool
	version    bool
	portNumber int
)

func printHelp() {
	flag.PrintDefaults()
	os.Exit(0)
}

func printVersion() {
	fmt.Println(versionNumber)
	os.Exit(0)
}

func printArgs() {
	if version {
		printVersion()
	}
	if help {
		printHelp()
	}
}

func init() {
	flag.BoolVar(&help, "help", false, "Prints out avaliable commands")
	flag.BoolVar(&version, "version", false, "prints the version number")
	flag.IntVar(&portNumber, "port", 8080, "set custom port number")
	flag.Parse()
}

func main() {
	printArgs()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := authapi.SetupRouter()

	roll.Token = os.Getenv("RollBarToken")
	roll.Endpoint = os.Getenv("RollBarEnv")
	//roll.Environment = "production" // defaults to "development"
	router.Use(rollbar.Recovery(true))
	roll.Info("Starting server", map[string]string{})

	// start the server
	router.Run(fmt.Sprintf(":%d", portNumber))
}
