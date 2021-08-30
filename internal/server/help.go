package server

import "github.com/gofiber/fiber/v2"

func help(c *fiber.Ctx) error {
	return c.SendString("Hello 👋👋👋👋👋👋👋👋👋👋👋👋👋👋!")
}
