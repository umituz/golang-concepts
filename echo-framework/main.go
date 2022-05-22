package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Everything is OK!")
}

func userHandler(c echo.Context) error {
	dataType := c.Param("dataType")
	username := c.QueryParam("username")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Username : %s Name : %s Surname : %s", username, name, surname))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"username": username,
			"name":     name,
			"surname":  surname,
		})
	}

	return c.String(http.StatusBadRequest, "Data type only should be json or string")
}

func main() {

	e := echo.New()

	e.GET("/", homeHandler)
	e.GET("/users/:dataType", userHandler)

	e.Start(":8000")

}
