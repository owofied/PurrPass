package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	
	"github.com/XiroXD/PurrPass/router"
)

func main() {


	app := fiber.New()

	app.Use(logger.New())

	router.Setup(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"statusCode": fiber.StatusNotFound,
			"error":      "Not Found",
			"message":    fmt.Sprintf("Cannot %s %s", c.Method(), c.Path()),
		})
	})

	app.Listen(":3000")
}
