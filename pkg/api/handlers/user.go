package handlers

import (
	"github.com/Vallghall/book-list/pkg/store"
	"github.com/Vallghall/book-list/pkg/types"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

// CreateUser - handles user creation
func CreateUser(db store.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		uc := new(types.UserCreation)
		err := c.BodyParser(uc)
		if err != nil {
			return fiber.ErrBadRequest
		}

		user, err := db.CreateUser(uc)
		if err != nil {
			return err // TODO: improve
		}

		return c.JSON(user)
	}
}

// GetUser - handles getting user by id
func GetUser(db store.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return ErrMissingIDParam
		}

		uid, err := uuid.FromString(id)
		if err != nil {
			return ErrBadUUID
		}

		user, err := db.GetUser(uid)
		if err != nil {
			return err // TODO: improve
		}

		return c.JSON(user)
	}
}
