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
	// server.Use(cors.Default())
	server.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
	}))
	routes.RegisterRoutes(server)
	// Start the server on port 8080
	server.Run(":8080")
}
