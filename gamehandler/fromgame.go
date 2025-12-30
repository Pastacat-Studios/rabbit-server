package gamehandler

import (
	"pastacat/rabbitserver/database"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Score struct {
	Id    string `json:"id" db:"id"`
	Score int    `json:"score" db:"score"`
}

func GetGameJson(c *gin.Context) {
	var newscore Score
	err := c.ShouldBindBodyWith(&newscore, binding.JSON)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err = database.DB.NamedExec(`INSERT INTO scores (id, score) VALUES (:id, :score)`, newscore)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Submitted Successfully",
	})
}
