package lfdlapi

import (
	"fmt"
	"log"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// noRouteHandler is a general no route handler
func noRouteHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}

//// healthCheckHandler is a general health check handler
//func healthCheckHandler(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"ServiceName": "xyz",
//		"Version":     "0.0.1",
//		"health":      "good",
//	})
//}
//
//func whoAmI(c *gin.Context) {
//	claims := jwt.ExtractClaims(c)
//	c.JSON(200, gin.H{
//		"userID": claims["id"],
//	})
//}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":       "pong",
		"systemTimeUTC": time.Now(),
	})
}

func createEventHandler(c *gin.Context) {
	startTime := c.Query("starttime")
	endTime := c.Query("endtime")
	points := c.Query("points")
	eventType := c.Query("type")
	severity := c.Query("severity")
	centerX := c.Query("centerx")
	centerY := c.Query("centery")

	startTimeConvert, err := time.Parse(time.RFC3339, startTime)
	endTimeConvert, err := time.Parse(time.RFC3339, endTime)


	if startTime =="" || endTime == "" ||   points == "" || eventType == "" || severity == "" || centerX == "" || centerY == "" || err != nil {

		message := ""
		//
		if startTime == "" {
			message = message + "Start time can not be empty"
		}
		if endTime == "" {
			message = message + ": End time can not be empty"
		}
		if points == "" {
			message = message + ": points can not be empty"
		}
		if eventType == "" {
			message = message + ": event type can not be empty"
		}
		if severity == "" {
			message = message + ": severity can not be empty"
		}
		if centerX == "" {
			message = message + ": center x can not be empty"
		}
		if centerY == "" {
			message = message + ": center y can not be empty"
		}

		if err != nil {
			message = message + ": error parsing startTime or endTime"
		}

		fmt.Println(message, err)

		c.JSON(418, gin.H{
			"message": message,
		})
	} else {


		create := db.Create(&Event{
			StartTime: startTimeConvert,
			EndTime:   endTimeConvert,
			Points:    points,
			EventType: eventType,
			Severity:  severity,
			CenterX: centerX,
			CenterY: centerY,
		})

		c.JSON(200, gin.H{
			"create": create.Value,
		})
	}
}

func readEventHandler(c *gin.Context) {
	topLatitude := c.Query("toplatitude")
	bottomLatitude := c.Query("bottomlatitude")
	leftLongitude := c.Query("leftlongitude")
	rightLongitude := c.Query("rightlongitude")

	if topLatitude == ""|| bottomLatitude == "" || leftLongitude == "" || rightLongitude == "" {
		message := ""

		if topLatitude == "" {
			message = message + ": topLatitude can not be empty"
		}
		if bottomLatitude == "" {
			message = message + ": bottomLatitude can not be empty"
		}
		if leftLongitude == "" {
			message = message + ": leftLongitude can not be empty"
		}
		if rightLongitude == "" {
			message = message + ": rightLongitude can not be empty"
		}

		c.JSON(218, gin.H{
			"message": "bad pram",
		})
	} else {

		topLatitudeInt , _ := strconv.ParseFloat(topLatitude, 64)
		bottomLatitudeInt, _ := strconv.ParseFloat(bottomLatitude,64)
		leftLongitudeInt, _ := strconv.ParseFloat(leftLongitude,64)
		rightLongitudeInt, _ := strconv.ParseFloat(rightLongitude,64)

		var events []Event
		db.Where("center_y BETWEEN ? AND ?",bottomLatitudeInt, topLatitudeInt).Where("center_x BETWEEN ? AND ?",leftLongitudeInt, rightLongitudeInt).Find(&events)

		c.JSON(200, gin.H{
			"events": events,
		})

	}

}

func deleteEventHandler(c *gin.Context) {

	id := c.Query("id")

	message := ""

	if id == "" {
		message = message + ": id can not be empty"
	}

	idInt, err := strconv.Atoi(id)

	if err != nil {
		message = message + "id must be a int"
	}

	db.Where("id = ?", idInt).Delete(&Event{})
		c.JSON(200, gin.H{
			"message": message,
		})
}
