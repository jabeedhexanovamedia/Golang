package handlers

import (
	"db-connection/models"
	"db-connection/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetUsers returns a list of users.
func GetUsers(client *mongo.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        users, err := services.FindUsers(client, nil)
        if err != nil {
            c.JSON(500, gin.H{"error": "Failed to fetch users"})
            return
        }
        c.JSON(200, users)
    }
}

// CreateUser creates a new user.
func CreateUser(client *mongo.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": "Invalid request"})
            return
        }
        result, err := services.InsertUser(client, user)
        if err != nil {
            c.JSON(500, gin.H{"error": "Failed to create user"})
            return
        }
        c.JSON(200, result)
    }
}