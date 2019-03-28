package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lookfirstdrivelater/lfdlapi/api/eventapi"
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
	flag.IntVar(&portNumber, "port", 8081, "set custom port number")
	flag.Parse()
}

func main() {
	printArgs()

	router := eventapi.SetupRouter()

	router.Run(fmt.Sprintf(":%d", portNumber))

}
