package eventapi

import (
	"time"

	"github.com/gin-gonic/gin"
)

func healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":       "systems ok",
		"systemTimeUTC": time.Now(),
	})
}
