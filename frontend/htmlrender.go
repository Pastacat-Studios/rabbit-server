package frontend

import (
	"github.com/gin-gonic/gin"
)

func GenLeaderboard(c *gin.Context) {
	c.HTML(200, "leaderboard.tmpl", gin.H{})
}
