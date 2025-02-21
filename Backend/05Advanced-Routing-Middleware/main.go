package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("Route Grouping")
	r := gin.Default()

	r.GET("/", func(c *gin.Context)	{

		c.JSON(200, gin.H{
			"message": "Route Group",
		})
	})
	
	// Grouping

	v1 := r.Group("/api/v1")
	{
		// http://localhost:8080/api/v1/users
		v1.GET("/users", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message":"Get all the users(v1)",
			})
		})

		// http://localhost:8080/api/v1/admin
		v1.GET("/admin", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message":"Get all the admin(v1)",
			});
		})
	}

	v2 := r.Group("/api/v2")
	{
		// http://localhost:8080/api/v2/users
		v2.GET("/users", func(c *gin.Context){
			c.JSON(200, gin.H{
				"message":"Get all the users(v2)",
			})
		})
	}

	r.Run(":8080")

	
}