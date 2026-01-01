package frontend

import (
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
	c.HTML(200, "user.tmpl", gin.H{
		"user": user.Id,
	})
}
