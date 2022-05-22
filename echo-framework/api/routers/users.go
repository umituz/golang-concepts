package routers

import (
	"github.com/labstack/echo/v4"
	"golangconcepts/echo-framework/api/handlers"
)

func UserRouter(g *echo.Group) {
	g.GET("/users/:dataType", func(c echo.Context) error {
		return handlers.UsersHandler(c)
	})

	g.POST("/users/add", func(c echo.Context) error {
		return handlers.UsersAddHandler(c)
	})

}
