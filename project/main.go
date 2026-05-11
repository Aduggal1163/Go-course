package main

import (
	"net/http"
	"strconv"

	"example.com/m/db"
	"example.com/m/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// Create a Gin router with default middleware (logger + recovery)
	server := gin.Default()
	server.Use(cors.Default())
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)
	// Start the server on port 8080
	server.Run(":8080")
}

func getEvents(context *gin.Context) { // context *gin.Context = request + response handler object. CORE THING INSIDE EVERY ROUTE
	// Send JSON response with status 200 (OK)
	// gin.H is just a shortcut for map[string]interface{}
	// context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events! Try again after sometime"})
		return
	}
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
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Save/create  events! Try again after sometime"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Created Successfully", "event:": event})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse id! Try again after sometime"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not fetch event from this id! Try again after sometime"})
		return
	}
	context.JSON(http.StatusOK, event)
}
