package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lookfirstdrivelater/lfdlapi/internal/lfdlapi"
	"log"
	"os"

	"github.com/joho/godotenv"

)

const versionNumber = "0.3.0"

var (
	version bool
	help    bool
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
	flag.Parse()
}

func main() {
	// print out argument
	printArgs()

	//load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	args := fmt.Sprint(os.Getenv("DatabaseUsername") + ":" + os.Getenv("DatabasePassword") + "@/" + os.Getenv("DatabaseName") + "?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", args)
	if err != nil {
		log.Println(err)
		log.Panic("database failed to open")
	}
	defer db.Close()

	// after other services are added then we can make this a go routene
	lfdlapi.API(db)
}
