package main

import (
	"html/template"
	"os"
	"pastacat/rabbitserver/database"
	"pastacat/rabbitserver/frontend"
	"pastacat/rabbitserver/gamehandler"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect(os.Getenv("DB_PATH"))
	router := gin.Default()
	api := router.Group("api")
	{
		api.GET("/maxscore", gamehandler.SendHighestScore)
		gameonly := api.Group("/", gamehandler.CheckIfGame, gamehandler.CheckUsername)
		{
			gameonly.POST("/connect", gamehandler.PongGame)
			gameonly.POST("/submit", gamehandler.GetGameJson)
		}
	}
	router.SetFuncMap(template.FuncMap{
		"listofscores": frontend.GenScoreList,
	})
	router.LoadHTMLGlob("frontend/templates/*.tmpl")
	router.GET("/leaderboard", frontend.GenLeaderboard)
	router.Run() // listens on 0.0.0.0:8080 by default
}
