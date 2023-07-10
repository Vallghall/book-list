package api

import "github.com/gofiber/fiber/v2"

// InitApp - initialize routes
func InitApp() *fiber.App {
	app := fiber.New(fiber.Config{
		PassLocalsToViews:     true,
		DisableStartupMessage: true,
		ErrorHandler:          fiber.DefaultErrorHandler,
		Views:                 engine(),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	return app
}
