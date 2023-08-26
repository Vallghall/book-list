package handlers

import (
	mw "github.com/Vallghall/book-list/pkg/middleware"
	"github.com/Vallghall/book-list/pkg/types"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

// CreateBook - book creation
func CreateBook(c *fiber.Ctx) error {
	bc := new(types.BookCreation)
	if err := c.BodyParser(bc); err != nil {
		return fiber.ErrBadRequest
	}

	repo := mw.Repo(c)
	uid := mw.UserID(c)

	book, err := repo.CreateBook(uid, bc)
	if err != nil {
		return err
	}

	return c.JSON(book)
}

// GetBookByID - fetching book info by its id
func GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := uuid.FromString(id)
	if err != nil {
		return fiber.ErrBadRequest
	}

	repo := mw.Repo(c)

	book, err := repo.GetBookByID(uid)
	if err != nil {
		return err
	}

	return c.JSON(book)
}
