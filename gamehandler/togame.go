package gamehandler

import (
	"pastacat/rabbitserver/database"
	"regexp"

	"github.com/gin-gonic/gin"
)

const Version string = "0.1a"

type connection struct {
	Ver string `json:"version"`
	Id  string `json:"id"`
}

func PongGame(c *gin.Context) {
	var con connection
	err := c.BindJSON(&con)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if con.Ver != Version {
		c.JSON(400, gin.H{"error": "Incompatible Version"})
		return
	}
	match, _ := regexp.MatchString(`^[a-zA-Z]+$`, con.Id)
	if !match {
		c.JSON(406, gin.H{"error": "Bad Username"})
		return
	}
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
		c.JSON(200, gin.H{"score": 0}) //Unimportant enough to silently fail
		return
	}
	c.JSON(200, gin.H{
		"score": score.Score,
	})
}
