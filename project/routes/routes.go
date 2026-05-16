package routes

import (
	"example.com/m/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", middlewares.Authentication, createEvents)
	server.PUT("/events/:id", middlewares.Authentication, updateEvent)
	server.DELETE("/events/:id", middlewares.Authentication, deleteEvent)
	server.POST("/signup", signup)
	server.POST("/signin", signin)
	server.POST("/events/:id/register", middlewares.Authentication, registerForEvent)
	server.DELETE("/events/:id/register", middlewares.Authentication, cancelRegisteration)
}
