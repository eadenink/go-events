package main

import (
	"net/http"
	"time"

	"github.com/eadenink/go-events/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"events": models.GetEvents(),
	})
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event.ID = len(models.GetEvents()) + 1
	event.DateTime = time.Now()
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"event": event,
	})
}
