package main

import (
	"fmt"
	"net/http"
	"strconv"

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

	// Path parameters (/user/:id)
	e.GET("/user/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, "user id: "+id)
	})

	// multiple param
	e.GET("/user/:userId/org/:orgId", func(c echo.Context) error {

		userId := c.Param("userId")
		orgId := c.Param("orgId")
		params := c.ParamValues()

		res := map[string]interface{}{
			"userId": userId,
			"orgId":  orgId,
			"params": params,
		}

		return c.JSON(200, res)
	})

	// validation and transformation
	e.GET("user/:userId/post/:postId", func(c echo.Context) error {

		//validation

		userId := c.Param("userId")
		postId := c.Param("postId")

		if userId == "" || postId == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "user id and post id required",
			})
		}

		// if userId <= 0 {
		// 	return c.String(http.StatusBadRequest, "user id must be postive number only")
		// }
		// transformation
		idVal, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"message": "user id mustbe valid id",
				"error":   err.Error(),
			})
		}

		if idVal <= 0 {
			return c.String(http.StatusBadRequest, "user id must be postive number only")
		}

		// Business logic
		return c.JSON(200, map[string]interface{}{
			"userId": idVal,
			"postId": postId,
		})

	})
	e.Logger.Fatal(e.Start("localhost:1323"))

}
