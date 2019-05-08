package housekeeper

import (
	"github.com/jinzhu/gorm"
	"github.com/lookfirstdrivelater/lfdlapi/internal/lfdlapi"
	"time"
)

// HouseKeeper
func HouseKeeper(DB *gorm.DB) {

	houseKeep(DB)

	ticker := halfHourTicker()
	for  range ticker.C{
		ticker = halfHourTicker()
		houseKeep(DB)
	}
}

func houseKeep(DB *gorm.DB) {
	cutOffTime := time.Now().Add(-time.Hour * time.Duration(24))

	var events []lfdlapi.Event
	DB.Find(&events)

	for i := 0; i < len(events); i++ {

		if events[i].EndTime.Before(cutOffTime) {
			DB.Where("id = ?", events[i].ID).Delete(lfdlapi.Event{})
		}
	}
}

func halfHourTicker() *time.Ticker {
	untilHalfHour := (time.Minute * time.Duration(30-time.Now().Minute())) + (time.Second * time.Duration(60-time.Now().Second())) - time.Duration(time.Minute*1)
	if untilHalfHour <= 0 {
		untilHalfHour = untilHalfHour + time.Duration(time.Minute*30)
	}

	return time.NewTicker(untilHalfHour)
}
