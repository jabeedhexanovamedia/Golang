package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// func main() {
// 	// fmt.Println("Hello, World!")

// 	e := echo.New()

// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Welcome to echoo!")
// 	})
// 	e.Start(":8080")
// }

func main() {
	e := echo.New()

	e.GET("/welcome", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Lcome",
		})
	})

	e.Logger.Fatal(e.Start(":1233"))
}
