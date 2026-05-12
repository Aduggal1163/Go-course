package main

import (
	"example.com/m/db"
	"example.com/m/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// Create a Gin router with default middleware (logger + recovery)
	server := gin.Default()
	server.Use(cors.Default())
	routes.RegisterRoutes(server)
	// Start the server on port 8080
	server.Run(":8080")
}
