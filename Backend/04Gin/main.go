// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
//     // Create a new Gin router
//     r := gin.Default()

//     // Define a route for GET /hello
//     r.GET("/", func(c *gin.Context) {
//         c.JSON(200, gin.H{
//             "message": "Hello, World!",
//         })
//     })

//     // Run the server on port 8080
//     r.Run(":8080")
// }

//! Homework for Day 1
// Write a Gin application with the following endpoints:

// GET /ping → Returns {"message": "pong"}.

// GET /user/:id → Returns {"user_id": "{id}"}.

// GET /search?q=term → Returns {"query": "term"}.

// Add a custom middleware that logs the request method and path.

// Experiment with c.Set() and c.Get() to pass data between middleware and handlers.

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Add a custom middleware that logs the request method and path.
func customMiddleware(c *gin.Context){
	fmt.Println("Request method and path: ", c.Request.Method, c.Request.URL.Path)
	c.Next()

}

// Experiment with c.Set() and c.Get() to pass data between middleware and handlers.
// - In Gin, c.Set() and c.Get() allow you to store and retrieve data within the same request context, which is useful for passing data between middleware and handlers.

func storeUserNameMiddleare (c *gin.Context) {

	c.Set("userName", "Jabeed")
	c.Next()
}

func printUserNameHandler (c *gin.Context) {

	// Retrieve the value set in the middleware

	userName, exists := c.Get("userName")

	if !exists {
		c.JSON(404, gin.H{
			"error":"user not found",
		})
	}

	c.JSON(200, gin.H{"message": "Hello, " + userName.(string)})
}

func main () {

	r:= gin.Default()
	// Add a custom middleware that logs the request method and path.
	r.Use(customMiddleware)

	// GET /ping → Returns {"message": "pong"}.
	r.GET("/ping", func (c *gin.Context){

		c.JSON(200, gin.H{
			"message":"Pong",
		})
	})

	// GET /user/:id → Returns {"user_id": "{id}"}.
	r.GET("/user/:id", func (c *gin.Context){

		id := c.Param("id")

		c.JSON(200, gin.H{
			"user_id":id,
		})
	})

	// GET /search?q=term → Returns {"query": "term"}.
	r.GET("/search", func(c *gin.Context) {
		term := c.Query("q")

		c.JSON(200, gin.H{
			"query": term,
		})
	})
// Add a custom middleware that logs the request method and path.
	r.Use(customMiddleware)

		// Apply middleware to a route
		r.GET("/welcome", storeUserNameMiddleare, printUserNameHandler)


	r.Run(":8080")
}