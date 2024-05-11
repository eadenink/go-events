package main

import (
	"github.com/eadenink/go-events/db"
	eventMethods "github.com/eadenink/go-events/methods/events"
	userMethods "github.com/eadenink/go-events/methods/users"
	"github.com/eadenink/go-events/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()

	server := gin.Default()

	events := server.Group("/events")

	//----- EVENT ROUTES -----

	// Unauthenticated

	events.GET("/", eventMethods.GetEvents)
	events.GET("/:id", eventMethods.GetEvent)

	// Authenticated
	events.Use(middlewares.CheckAuth)

	events.POST("/", eventMethods.CreateEvent)
	events.PUT("/:id", eventMethods.UpdateEvent)
	events.DELETE("/:id", eventMethods.DeleteEvent)

	//----- USER ROUTES -----

	server.POST("/signup", userMethods.SignUp)
	server.POST("/login", userMethods.Login)

	server.Run(":8080")
}
