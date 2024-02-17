package main

import (
	"github.com/labstack/echo/v4"
	"github.com/owofied/PurrPass/router"
)

func main() {
	e := echo.New()

	router.SetupRouter(e)

	e.Logger.Fatal(e.Start(":3000"))
}
