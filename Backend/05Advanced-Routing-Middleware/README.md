### Advanced Routing, Middleware, and Request Handling:

##### 1. Advanced Routing in Gin

##### What is Routing?

- Routing is the process of mapping incoming HTTP requests to the appropriate handler function based on the request path and HTTP method (GET, POST, etc.).
- In Gin, routing is done using the r.GET(), r.POST(), etc., methods.

##### Route Parameters

- Route parameters allow you to capture dynamic values from the URL path.
- Example:

```go

r.GET("/user/:id", func(c *gin.Context) {
    id := c.Param("id") // Capture the "id" parameter
    c.JSON(200, gin.H{
        "user_id": id,
    })
})

```

- A request to /user/123 will return {"user_id": "123"}.

##### How It Works:

- The :id syntax tells Gin to treat id as a dynamic parameter.

- Gin stores the captured value in c.Params, which is a slice of gin.Param structs.

- You can access the value using c.Param("id").

##### Query Parameters

- Query parameters are key-value pairs in the URL after the ?.
- Example:

```go

r.GET("/search", func(c *gin.Context) {
    query := c.Query("q") // Capture the "q" query parameter
    c.JSON(200, gin.H{
        "query": query,
    })
})

```

- A request to /search?q=golang will return {"query": "golang"}.

##### How It Works:

- Gin parses the query string using url.ParseQuery() and stores the results in c.Request.URL.Query().
- The c.Query() method retrieves the value of a specific query parameter.

#### Route Grouping in Gin

##### What is Route Grouping?

- Route grouping is a way to organize your routes into logical groups, often based on a common path prefix or functionality.

- It helps you:
- Avoid repeating the same path prefix for multiple routes.
- Apply middleware to a group of routes.
- Keep your code clean and maintainable.

##### How Does Route Grouping Work?

- In Gin, you create a route group using the Group() method.
- Example

```go
api := r.Group("/api")
{
   api.GET("/users", func(c *gin.Context) {
       c.JSON(200, gin.H{
           "message": "Get all users",
       })
   })
   api.POST("/users", func(c *gin.Context) {
       c.JSON(200, gin.H{
           "message": "Create a new user",
       })
   })
}
```

- Here, /api is the common prefix for all routes in the group.
- The routes will be:
- GET /api/users
- POST /api/users

##### Real-World Example: API Versioning

- In real-world applications, APIs often have multiple versions (e.g., v1, v2).
- Route grouping makes it easy to handle versioning.
- Example:

```go
v1 := r.Group("/api/v1")
{
    v1.GET("/users", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Get all users (v1)",
        })
    })
}

v2 := r.Group("/api/v2")
{
    v2.GET("/users", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Get all users (v2)",
        })
    })
}
```

- This creates routes like:
- GET /api/v1/users
- GET /api/v2/users

##### Applying Middleware to Route Groups

-Middleware can be applied to a specific route group.

- Example:

```go
auth := r.Group("/auth")
auth.Use(AuthMiddleware) // Apply authentication middleware to this group
{
    auth.GET("/profile", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "User profile",
        })
    })
}
```

- Here, the AuthMiddleware will be applied to all routes in the /auth group.
- Group(): Creates a new route group with a common prefix.
- Use(): Adds middleware to the route group.

##### 2. Advanced Middleware in Gin

- Middleware is a function that processes a request before it reaches the handler.

- It can be used for:
- Logging requests.
- Authentication and authorization.
- Request validation.
- Error handling.
- Adding custom headers.

##### How Does Middleware Work?

- Middleware functions are executed in the order they are added.

- Each middleware function has access to the \*gin.Context object, which encapsulates the request and response.

Middleware can:

- Modify the request (e.g., add data to the context).
- Abort the request (e.g., if authentication fails).
- Pass control to the next middleware or handler using c.Next().

##### 1. Real-World Example: Logging Middleware

- Let’s create a middleware that logs the request method, path, and response time.
- Example:

```go
func LoggerMiddleware(c *gin.Context) {
    // Code to execute before the handler
    start := time.Now() // Record the start time

    c.Next() // Pass control to the next middleware or handler

    // Code to execute after the handler
    duration := time.Since(start) // Calculate the response time
    fmt.Printf("Request: %s %s | Response Time: %v\n", c.Request.Method, c.Request.URL.Path, duration)
}

```

- Add the Middleware to the Router:

```go

// Add the Middleware to the Router:
r := gin.Default()
r.Use(LoggerMiddleware) // Add the logging middleware

r.GET("/hello", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello, World!",
    })
})

r.Run(":8080")
```

##### How It Works:

- When a request is made, the LoggerMiddleware is executed first.
- It records the start time, then calls c.Next() to pass control to the next middleware or handler.
- After the handler completes, the middleware calculates the response time and logs the details.

##### 2. Real-World Example: Authentication Middleware

- Let’s create a middleware that checks for a valid API key in the request headers.
- Example:

```go

func AuthMiddleware(c *gin.Context) {
    apiKey := c.GetHeader("Authorization")
    if apiKey != "my-secret-key" {
        c.JSON(401, gin.H{
            "error": "Unauthorized",
        })
        c.Abort() // Stop the request from proceeding further
        return
    }
    c.Next() // Pass control to the next middleware or handler
}


```

- Add the Middleware to a Route Group:
- Apply the middleware to a specific route group.

```go
auth := r.Group("/admin")
auth.Use(AuthMiddleware) // Apply authentication middleware to this group
{
    auth.GET("/dashboard", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Admin Dashboard",
        })
    })
}
```

##### How It Works:

- When a request is made to /admin/dashboard, the AuthMiddleware is executed first.
- If the Authorization header is missing or invalid, the middleware sends a 401 Unauthorized response and aborts the request.
- If the API key is valid, the middleware calls c.Next() to pass control to the handler.

##### 3. Real-World Example: Error Handling Middleware

- Let’s create a middleware that catches any panics and returns a 500 Internal Server Error response.
- The middleware uses defer and recover() to catch panics.
- Example:

```go
func RecoveryMiddleware(c *gin.Context) {
    defer func() {
        if err := recover(); err != nil {
            c.JSON(500, gin.H{
                "error": "Internal Server Error",
            })
            c.Abort() // Stop the request from proceeding further
        }
    }()
    c.Next() // Pass control to the next middleware or handler
}
```

- Add the Middleware to the Router:

```go
r := gin.New()
r.Use(RecoveryMiddleware) // Add the recovery middleware

r.GET("/panic", func(c *gin.Context) {
    panic("Something went wrong!") // Simulate a panic
})

r.Run(":8080")
```

##### How It Works:

- When a request is made to /panic, the handler intentionally panics.
- The RecoveryMiddleware catches the panic, sends a 500 Internal Server Error response, and aborts the request.

##### 4. Real-World Example: Custom Headers Middleware

- Let’s create a middleware that adds custom headers to the response.
- The middleware adds a X-Custom-Header to the response.

```go
func CustomHeaderMiddleware(c *gin.Context) {
    c.Header("X-Custom-Header", "Hello from Middleware!")
    c.Next() // Pass control to the next middleware or handler
}
```

- Add the Middleware to the Router:

```go
r := gin.Default()
r.Use(CustomHeaderMiddleware) // Add the custom header middleware

r.GET("/hello", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello, World!",
    })
})

r.Run(":8080")
```

##### How It Works:

- When a request is made to /hello, the CustomHeaderMiddleware adds the X-Custom-Header to the response.
- The handler then processes the request and sends the response.

##### Keywords

- c.Next(): Passes control to the next middleware or handler.
- c.Abort(): Stops the request from proceeding further.
- c.GetHeader(): Retrieves the value of a request header.
- c.Header(): Sets a response header.
- defer: Schedules a function to run when the surrounding function returns.
- recover(): Catches a panic and allows the program to continue running.

##### Practice

- Create a middleware that validates the Content-Type header and ensures it is application/json.

- Add a middleware that limits the size of incoming requests to 1 MB
