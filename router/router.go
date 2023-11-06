package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XiroXD/PurrPass/controllers"
)

func Setup(app *fiber.App) {
	baseController := controllers.NewBaseController()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", baseController.Hello)
}
