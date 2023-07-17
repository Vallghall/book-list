package handlers

import (
	"github.com/Vallghall/book-list/pkg/store"
	"github.com/Vallghall/book-list/pkg/types"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

// CreateAuthor - author creation handler
func CreateAuthor(db store.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ac := new(types.AuthorCreation)
		err := c.BodyParser(ac)
		if err != nil {
			return fiber.ErrBadRequest
		}

		// FIXME: fix it where user auth and whatnot will be implemented
		author, err := db.CreateAuthor(uuid.Nil, ac)
		if err != nil {
			return err
		}

		return c.JSON(author)
	}
}

// GetAuthorByID - author getting by id
func GetAuthorByID(db store.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return ErrMissingIDParam
		}

		uid, err := uuid.FromString(id)
		if err != nil {
			return ErrBadUUID
		}

		author, err := db.GetAuthorByID(uid)
		if err != nil {
			return err // TODO: improve
		}

		return c.JSON(author)
	}
}
