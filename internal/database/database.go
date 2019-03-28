package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Users struct {
	gorm.Model
	UserName  string
	FirstName string
	LastName  string
}

type Event struct {
	gorm.Model
	StartTime time.Time
	EndTime   time.Time
}

// Connect create a database connection
func connect(user, password, dbname string) (*gorm.DB, error) {
	args := fmt.Sprint(user + ":" + password + "@/" + dbname + "?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", args)
	defer db.Close()
	return db, err
}
