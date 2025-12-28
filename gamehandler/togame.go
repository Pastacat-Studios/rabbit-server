package gamehandler

import "github.com/gin-gonic/gin"

func PongGame(c *gin.Context) {
	c.JSON(200, gin.H{
		"server":  "Rabbit Server",
		"version": "0.1a",
	})
}
