package store

import (
	"github.com/Vallghall/book-list/pkg/types"
	uuid "github.com/satori/go.uuid"
)

// Store - contract for types that serve
// as a database interface
type Store interface {
	CreateUser(c *types.UserCreation) (*types.User, error)
	GetUser(userID uuid.UUID) (*types.User, error)

	CreateAuthor(userID uuid.UUID, c *types.AuthorCreation) (*types.Author, error)
	GetAuthorByID(authorID uuid.UUID) (*types.Author, error)

	CreateBook(c *types.BookCreation) (*types.Book, error)
	GetBookByID(bookID uuid.UUID) (*types.Book, error)
}