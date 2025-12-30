package gamehandler

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CheckIfGame(c *gin.Context) {
	check := c.GetHeader("Ring")
	if check != "Rabbit" {
		c.AbortWithStatus(401)
		return
	}
	c.Next()
}

type justid struct {
	Id string `json:"id"`
}

func CheckUsername(c *gin.Context) {
	var id justid
	err := c.ShouldBindBodyWith(&id, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "No username provided"})
		return
	}
	match, _ := regexp.MatchString(`^[a-zA-Z]+$`, id.Id)
	if !match {
		c.AbortWithStatusJSON(406, gin.H{"error": "Bad Username"})
		return
	}
	c.Next()
}
