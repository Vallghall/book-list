package middleware

import (
	"github.com/Vallghall/book-list/pkg/store"
	"github.com/gofiber/fiber/v2"
)

// WithDB - middleware decorator for adding
// Repository object into context's locals
func WithDB(r store.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("repo", r)
		return c.Next()
	}
}
