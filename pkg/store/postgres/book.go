package postgres

import (
	"github.com/Vallghall/book-list/pkg/types"
	uuid "github.com/satori/go.uuid"
)

// CreateBook - book creation
func (s *DB) CreateBook(c *types.BookCreation) (*types.Book, error) {
	book := new(types.Book)
	err := s.db.Get(book, "select * from library.create_book($1::uuid, $2::uuid, ($3,$4)::library.book)")

	return book, err
}

// GetBookByID - getting book by its ID
func (s *DB) GetBookByID(bookID uuid.UUID) (*types.Book, error) {
	book := new(types.Book)
	err := s.db.Get(book, "select * from library.get_book_by_id($1::uuid)")

	return book, err
}
