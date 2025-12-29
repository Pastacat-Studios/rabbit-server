package main

import (
	"pastacat/rabbitserver/database"
	"pastacat/rabbitserver/gamehandler"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	router := gin.Default()
	api := router.Group("api")
	{
		api.GET("/maxscore", gamehandler.SendHighestScore)
		gameonly := api.Group("/", gamehandler.CheckIfGame)
		{
			gameonly.POST("/connect", gamehandler.PongGame)
			gameonly.POST("/submit", gamehandler.GetGameJson)
		}
	}
	router.Run() // listens on 0.0.0.0:8080 by default
}
