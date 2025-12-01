package main

import (
	"db-connection/db"
	"db-connection/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	client, err := db.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseMongoDB(client)

	// Create Gin router
	r := gin.Default()

	// Define routes
	r.GET("/users", handlers.GetUsers(client))
	r.POST("/users", handlers.CreateUser(client))

	// Start the server
	r.Run(":8080")
}