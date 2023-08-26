package api

import (
	"github.com/Vallghall/book-list/configs"
	"github.com/Vallghall/book-list/pkg/api/handlers"
	mw "github.com/Vallghall/book-list/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

// InitApp - initialize routes
func InitApp(c *configs.Conf) *fiber.App {
	repo := c.Store()
	app := fiber.New(fiber.Config{
		PassLocalsToViews: true,
		ErrorHandler:      fiber.DefaultErrorHandler,
		Views:             engine(),
	})

	app.Use(
		mw.WithDB(repo),
		mw.WithLogger(c.HandlerLogLevel()),
	)

	app.Get("/", func(c *fiber.Ctx) error {
		// TODO: implement a homepage
		return c.SendString("Hello world!")
	})

	auth := app.Group("/auth")
	{
		// POST /auth/user - user creation
		auth.Post("/user", handlers.CreateUser)
		// GET /auth/user/:id
		auth.Get("/user/:id", handlers.GetUser)
	}

	// service logic endpoints
	service := app.Group("/service")
	{
		// author related endpoints
		authors := service.Group("/authors")
		{
			// POST /service/authors/ - author creation
			authors.Post("/", handlers.CreateAuthor)
			// GET /service/authors/:id
			authors.Get("/:id", handlers.GetAuthorByID)
		}

		// book related endpoints
		books := service.Group("/books")
		{
			// POST /service/books/ - author creation
			books.Post("/", handlers.CreateBook)
			// GET /service/books/:id
			books.Get("/:id", handlers.GetBookByID)
		}
	}

	return app
}
