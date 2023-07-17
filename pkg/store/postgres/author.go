package postgres

import (
	"github.com/Vallghall/book-list/pkg/types"
	uuid "github.com/satori/go.uuid"
)

// CreateAuthor - author creation
func (s *DB) CreateAuthor(userID uuid.UUID, c *types.AuthorCreation) (*types.Author, error) {
	au := new(types.Author)
	err := s.db.Get(au,
		"select * from library.create_author($1::uuid, ($2,$3,$4,$5)::library.author)",
		c.Name, c.LastName, c.FatherName, c.Birthday,
	)

	return au, err
}

// GetAuthorByID - getting author by id
func (s *DB) GetAuthorByID(authorID uuid.UUID) (*types.Author, error) {
	au := new(types.Author)
	err := s.db.Get(au, "select * from library.get_author_by_id($1::uuid")

	return au, err
}
