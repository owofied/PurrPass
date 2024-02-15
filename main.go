package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/owofied/PurrPass/database"
	"github.com/owofied/PurrPass/router"
	"github.com/owofied/PurrPass/utils"
)

func init() {
	godotenv.Load()
	utils.CheckRequiredVariables()

	database.ConnectPostgres()
}

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

	app.Listen(":" + os.Getenv("PORT"))
}
