package middleware

import (
	"github.com/Vallghall/book-list/pkg/store"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// UserID - retrieves authenticated user's ID
// from context's locals
func UserID(ctx *fiber.Ctx) uuid.UUID {
	return ctx.Locals("user-id").(uuid.UUID)
}

// Repo - retrieving Repository handle from
// context's locals
func Repo(ctx *fiber.Ctx) store.Store {
	return ctx.Locals("repo").(store.Store)
}

// Logger - retrieving logger from context's logger
func Logger(ctx *fiber.Ctx) *zap.Logger {
	return ctx.Locals("logger").(*zap.Logger)
}
