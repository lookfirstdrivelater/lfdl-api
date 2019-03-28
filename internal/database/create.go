package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

func CreateUser(db *gorm.DB, Username, Firstname, Lastname string) {
	db.Create(&Users{UserName: Username, FirstName: Firstname, LastName: Lastname})
}

func CreateEvent(db *gorm.DB, startTime, endTime time.Time) {
	db.Create(&Event{StartTime: startTime, EndTime: endTime})
}
