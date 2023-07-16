package api

import (
	"github.com/Vallghall/book-list/configs"
	"github.com/gofiber/fiber/v2"
)

// InitApp - initialize routes
func InitApp(c *configs.Conf) *fiber.App {
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
