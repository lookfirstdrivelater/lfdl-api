package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/rollbar"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/stvp/roll"
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

	if err != nil {
		log.Fatalln("database init failed")
	}

	args := fmt.Sprint("matt" + ":" + "Behnke22" + "@/" + "lfdlstaging" + "?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", args)
	if err != nil {
		log.Println(err)
		log.Panic("database failed to open")
	}
	defer db.Close()

	if !db.HasTable(Event{}) {
		db.AutoMigrate(Event{})
		log.Println("Migrating Event")
	}
	if !db.HasTable(Users{}) {
		db.AutoMigrate(Users{})
		log.Println("Migrating Users")

		pass, _ := HashPassword("Admin")

		db.Create(&Users{UserName: "Admin", FirstName: "Matt", LastName: "Behnke", Password: pass})
	}

	router := SetupRouter(db)

	roll.Token = os.Getenv("RollBarToken")
	roll.Endpoint = os.Getenv("RollBarEnv")
	//roll.Environment = "production" // defaults to "development"
	router.Use(rollbar.Recovery(true))
	roll.Info("Starting server", map[string]string{})

	// start the server
	router.Run(fmt.Sprintf(":%d", portNumber))
}
