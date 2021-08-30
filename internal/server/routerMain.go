package server

import "github.com/gofiber/fiber/v2"

func routerMain(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/help", help)
	app.Get("/ff", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	// Create a custom error with HTTP code 782
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message")
	})
}
