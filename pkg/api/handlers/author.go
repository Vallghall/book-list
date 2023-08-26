package handlers

import (
	"fmt"

	mw "github.com/Vallghall/book-list/pkg/middleware"
	"github.com/Vallghall/book-list/pkg/types"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

// CreateAuthor - author creation handler
func CreateAuthor(c *fiber.Ctx) error {
	ac := new(types.AuthorCreation)
	err := c.BodyParser(ac)
	if err != nil {
		return fmt.Errorf("parse body: %w", err)
	}
	repo := mw.Repo(c)

	// FIXME: fix it where user auth and whatnot will be implemented
	uid, _ := uuid.FromString("65dd0ff9-f75c-42d9-8f69-6069c86186eb")
	author, err := repo.CreateAuthor(uid, ac)
	if err != nil {
		return err
	}

	return c.JSON(author)
}

// GetAuthorByID - author getting by id
func GetAuthorByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return ErrMissingIDParam
	}

	uid, err := uuid.FromString(id)
	if err != nil {
		return ErrBadUUID
	}

	repo := mw.Repo(c)

	author, err := repo.GetAuthorByID(uid)
	if err != nil {
		return err // TODO: improve
	}

	return c.JSON(author)
}
