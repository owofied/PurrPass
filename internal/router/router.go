package router

import "github.com/labstack/echo/v4"

func SetupRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
}
