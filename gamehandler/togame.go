package gamehandler

import (
	"pastacat/rabbitserver/database"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const Version string = "2.2"

type connection struct {
	Ver string `json:"version"`
	Id  string `json:"id"`
}

func PongGame(c *gin.Context) {
	var con connection
	err := c.ShouldBindBodyWith(&con, binding.JSON)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if con.Ver != Version {
		c.JSON(400, gin.H{"error": "Incompatible Version"})
		return
	}
	c.JSON(200, gin.H{
		"server":  "Rabbit Server",
		"version": Version,
	})
}

type tinyscore struct {
	Score int `db:"MAX(score)"`
}

func SendHighestScore(c *gin.Context) {
	score := tinyscore{}
	err := database.DB.Get(&score, "SELECT MAX(score) FROM scores")
	if err != nil {
		c.JSON(200, gin.H{"score": 0}) //Unimportant enough to silently fail
		return
	}
	c.JSON(200, gin.H{
		"score": score.Score,
	})
}
