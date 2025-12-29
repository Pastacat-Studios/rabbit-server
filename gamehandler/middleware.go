package gamehandler

import (
	"github.com/gin-gonic/gin"
)

func CheckIfGame(c *gin.Context) {
	check := c.GetHeader("Ring")
	if check != "Rabbit" {
		c.AbortWithStatus(401)
		return
	}
	c.Next()
}
