package main

import (
	"fmt"
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

// func main() {
// 	e := echo.New()

// 	e.GET("/welcome", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, map[string]string{
// 			"message": "Lcome",
// 		})
// 	})

// 	e.Logger.Fatal(e.Start(":1323"))
// }

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	e.GET("/json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"data": "json data",
		})
	})

	e.GET("/req", func(c echo.Context) error {

		// res := map[string]interface{}{
		// 	"method":  c.Request().Method,
		// 	"host":    c.Request().Host,
		// 	"url":     c.Request().URL.String(),
		// 	"uri":     c.Request().RequestURI,
		// 	"proto":   c.Request().Proto,
		// 	"headers": c.Request().Header,
		// }
		// return c.JSON(200, res)
		//
		return c.String(200, fmt.Sprintf("%+v", c.Request()))
	})

	e.Logger.Fatal(e.Start("localhost:8080"))

}
