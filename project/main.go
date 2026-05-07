package main

import (
	"net/http"
	"example.com/m/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger + recovery)
	server := gin.Default()
	server.Use(cors.Default())
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	// Start the server on port 8080
	server.Run(":8080")
}

func getEvents(context *gin.Context) { // context *gin.Context = request + response handler object. CORE THING INSIDE EVERY ROUTE
	// Send JSON response with status 200 (OK)
	// gin.H is just a shortcut for map[string]interface{}
	// context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.BindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Created Successfully", "event:": event})
}