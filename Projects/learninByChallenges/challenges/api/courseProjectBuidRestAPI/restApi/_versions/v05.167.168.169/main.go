package main

import (
	"net/http"
	"restapi/db"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/ping", getPong)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8910") // localhost:8910
}

func getPong(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message:": "Pong!!!"})
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}