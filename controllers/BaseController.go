package controllers

import "github.com/gofiber/fiber/v2"

type BaseController struct{}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (c *BaseController) Hello(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}
