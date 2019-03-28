package database

import (
	"github.com/jinzhu/gorm"
)

func migrateUsers(db *gorm.DB) {
	db.AutoMigrate(&Users{})
}

func migrateEvent(db *gorm.DB) {
	db.AutoMigrate(&Event{})
}
