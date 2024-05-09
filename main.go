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
	server.GET("/events/:id", methods.GetEvent)

	server.POST("/events", methods.CreateEvent)
	server.PUT("/events/:id", methods.UpdateEvent)
	server.DELETE("/events/:id", methods.DeleteEvent)

	server.Run(":8080")
}
