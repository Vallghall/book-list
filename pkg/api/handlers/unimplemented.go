package handlers

import "github.com/gofiber/fiber/v2"

// NotImplemented - handler filler for endpoints
// that don't have their handlers implemented
func NotImplemented(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}
