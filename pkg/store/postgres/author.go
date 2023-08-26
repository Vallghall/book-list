package postgres

import (
	"github.com/Vallghall/book-list/pkg/models"
	"github.com/Vallghall/book-list/pkg/types"
	uuid "github.com/satori/go.uuid"
)

// CreateAuthor - author creation
func (s *DB) CreateAuthor(userID uuid.UUID, c *types.AuthorCreation) (*types.Author, error) {
	author := &models.Author{
		Name:        c.Name,
		LastName:    c.LastName,
		FatherName:  c.FatherName.NullString,
		Birthday:    c.Birthday.NullTime,
		CreatedByID: userID,
	}

	err := s.db.Create(author).Error
	if err != nil {
		return nil, err
	}

	return author.ToTarget(), nil
}

// GetAuthorByID - getting author by id
func (s *DB) GetAuthorByID(authorID uuid.UUID) (*types.Author, error) {
	author := &models.Author{ID: authorID}
	err := s.db.First(author).Error
	if err != nil {
		return nil, err
	}

	return author.ToTarget(), nil
}
