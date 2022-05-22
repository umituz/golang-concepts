package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func UsersHandler(c echo.Context) error {

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

func UsersAddHandler(c echo.Context) error {
	user := User{}

	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user)

	if err != nil {
		return err
	}

	fmt.Println(user)
	return c.JSON(http.StatusOK, user)
}
