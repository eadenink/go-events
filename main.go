package main

import (
	"github.com/eadenink/go-events/db"
	eventMethods "github.com/eadenink/go-events/methods/events"
	userMethods "github.com/eadenink/go-events/methods/users"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()

	server := gin.Default()

	server.GET("/events", eventMethods.GetEvents)
	server.GET("/events/:id", eventMethods.GetEvent)

	server.POST("/events", eventMethods.CreateEvent)
	server.PUT("/events/:id", eventMethods.UpdateEvent)
	server.DELETE("/events/:id", eventMethods.DeleteEvent)

	server.POST("/signup", userMethods.SignUp)

	server.Run(":8080")
}
