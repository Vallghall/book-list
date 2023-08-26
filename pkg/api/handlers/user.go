package handlers

import (
	mw "github.com/Vallghall/book-list/pkg/middleware"
	"github.com/Vallghall/book-list/pkg/types"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

// CreateUser - handles user creation
func CreateUser(c *fiber.Ctx) error {
	uc := new(types.UserCreation)
	err := c.BodyParser(uc)
	if err != nil {
		return fiber.ErrBadRequest
	}

	repo := mw.Repo(c)

	user, err := repo.CreateUser(uc)
	if err != nil {
		return err // TODO: improve
	}

	return c.JSON(user)
}

// GetUser - handles getting user by id
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return ErrMissingIDParam
	}

	uid, err := uuid.FromString(id)
	if err != nil {
		return ErrBadUUID
	}

	repo := mw.Repo(c)

	user, err := repo.GetUser(uid)
	if err != nil {
		return err // TODO: improve
	}

	return c.JSON(user)
}
