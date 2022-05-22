package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Everything is OK!")
}

func New() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", homeHandler)

	g := e.Group("/api")

	UserRouter(g)

	return e
}
