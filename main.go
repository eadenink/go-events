package main

import (
	"github.com/eadenink/go-events/db"
	methods "github.com/eadenink/go-events/methods/events"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()

	server := gin.Default()

	server.GET("/events", methods.GetEvents)
	server.POST("/events", methods.CreateEvent)

	server.Run(":8080")
}
