package routes

import (
	"net/http"
	"strconv"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

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

	userId := context.GetInt64("userId")
	event.UserID = userId
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

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse id! Try again after sometime"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch event from this id! Try again after sometime"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var updatedEvent models.Event
	err = context.BindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event! Try again after sometime"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated Successfully!"})
}
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not fetch event from this id! Try again after sometime"})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event! Try again after sometime", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted Successfully!"})
}
