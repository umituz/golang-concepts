package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Everything is OK!")
}

func setCustomHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		contentType := c.Request().Header.Get("Content-Type")

		if contentType != "application/json" {
			return c.String(http.StatusBadRequest, "We allow just JSON Request")
		}

		return next(c)
	}
}

func New() *echo.Echo {
	e := echo.New()
	e.Use(setCustomHeader)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "statusCode:${status}\n",
	}))
	e.Use(middleware.Recover())

	e.GET("/", homeHandler)

	g := e.Group("/api")

	UserRouter(g)

	return e
}
