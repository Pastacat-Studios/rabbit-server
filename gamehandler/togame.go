package gamehandler

import (
	"pastacat/rabbitserver/database"

	"github.com/gin-gonic/gin"
)

func PongGame(c *gin.Context) {
	c.JSON(200, gin.H{
		"server":  "Rabbit Server",
		"version": "0.1a",
	})
}

type tinyscore struct {
	Score int `db:"MAX(score)"`
}

func SendHighestScore(c *gin.Context) {
	score := tinyscore{}
	err := database.DB.Get(&score, "SELECT MAX(score) FROM scores")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"score": score.Score,
	})
}
