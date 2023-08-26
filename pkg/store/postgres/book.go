package postgres

import (
	"github.com/Vallghall/book-list/pkg/models"
	"github.com/Vallghall/book-list/pkg/types"
	uuid "github.com/satori/go.uuid"
)

// CreateBook - book creation
func (s *DB) CreateBook(userID uuid.UUID, c *types.BookCreation) (*types.Book, error) {
	book := &models.Book{
		Title:       c.Title,
		ReleaseYear: c.ReleaseYear.NullInt64,
		CreatedByID: userID,
		Authors:     make([]*models.Author, len(c.AuthorIDs)),
	}

	for i, a := range c.AuthorIDs {
		book.Authors[i].ID = a
	}

	err := s.db.Create(book).Error
	if err != nil {
		return nil, err
	}

	err = s.db.Model(book).Preload("Authors").Find(book).Error
	if err != nil {
		return nil, err
	}

	return book.ToTarget(), nil
}

// GetBookByID - getting book by its ID
func (s *DB) GetBookByID(bookID uuid.UUID) (*types.Book, error) {
	book := &models.Book{ID: bookID}
	err := s.db.Preload("Authors").First(book).Error
	if err != nil {
		return nil, err
	}

	return book.ToTarget(), err
}
