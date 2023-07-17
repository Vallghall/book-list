package api

import (
	"github.com/Vallghall/book-list/configs"
	"github.com/Vallghall/book-list/pkg/api/handlers"
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
	db := c.Store()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	auth := app.Group("/auth")
	{
		// POST /auth/user - user creation
		auth.Post("/user", handlers.CreateUser(db))
		// GET /auth/user/:id
		auth.Get("/user/:id", handlers.GetUser(db))
	}

	author := app.Group("/author")
	{
		// POST /author/ - author creation
		author.Post("/", handlers.CreateAuthor(db))
		// GET /author/:id
		author.Get("/user/:id", handlers.GetAuthorByID(db))
	}

	return app
}
