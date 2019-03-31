package lfdlapi

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	// there is a blank import
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type event struct {
	gorm.Model
	StartTime time.Time
	EndTime   time.Time
}

type user struct {
	gorm.Model
	UserName  string
	FirstName string
	LastName  string
	Password  string
	AuthGeneral bool
	AuthAdmin bool
}

// db ...
var db *gorm.DB

// API is the main entry point for this lib
func API(DB *gorm.DB) {

	db = DB

	// run migrations only if the tables do not exist
	if !DB.HasTable(event{}) {
		DB.AutoMigrate(event{})
		log.Println("Migrating Event")
	}
	if !DB.HasTable(user{}) {
		DB.AutoMigrate(user{})
		log.Println("Migrating Users")

		pass, err := hashPassword("admin")

		fmt.Println(err)

		db.Create(&user{UserName: "admin", FirstName: "Matt", LastName: "Behnke", Password: pass})
	}

	// create restfull router
	router := setupRouter()
	// start the server
	router.Run(fmt.Sprintf(":8080"))
}
