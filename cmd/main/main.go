package main

import (
	"errors"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/owofied/PurrPass/internal/config"
	"github.com/owofied/PurrPass/internal/router"
)

func main() {
	e := echo.New()

	confDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	if _, err = os.Stat(confDir + "/PurrPass/config.json"); errors.Is(err, os.ErrNotExist) {
		config.InitializeConfig()
	}

	conf := config.GetConfig()

	router.SetupRouter(e)

	e.Logger.Fatal(e.Start(":" + conf.Server.Port))
}
