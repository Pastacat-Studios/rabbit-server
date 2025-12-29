package main

import (
	"pastacat/rabbitserver/database"
	"pastacat/rabbitserver/gamehandler"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	router := gin.Default()
	router.POST("/api/connect", gamehandler.PongGame)
	router.GET("/api/maxscore", gamehandler.SendHighestScore)
	router.POST("/api/submit", gamehandler.GetGameJson)
	router.Run() // listens on 0.0.0.0:8080 by default
}
