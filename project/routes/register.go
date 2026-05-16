package routes

import (
	"net/http"
	"strconv"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id"})
		return
	}
	event, err := models.GetEventById(eventid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not find event by id"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event registered"})
}

func cancelRegisteration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse id"})
		return
	}
	event, err := models.GetEventById(eventid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not find event by id"})
		return
	}
	err = event.Cancel(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel user for event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event cancelled"})

}
