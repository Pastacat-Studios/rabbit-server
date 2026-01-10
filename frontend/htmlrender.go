package frontend

import (
	"regexp"

	"github.com/gin-gonic/gin"
)

func GenLeaderboard(c *gin.Context) {
	c.HTML(200, "leaderboard.tmpl", gin.H{})
}

type id struct {
	Id string `uri:"id"`
}

func GenUser(c *gin.Context) {
	var user id
	c.BindUri(&user)
	match, _ := regexp.MatchString(`^[a-zA-Z]+$`, user.Id)
	if !match {
		c.AbortWithStatusJSON(406, gin.H{"error": "Bad Username"})
		return
	}
	c.HTML(200, "user.tmpl", gin.H{
		"user": user.Id,
	})
}
