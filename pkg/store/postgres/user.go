package postgres

import (
	"github.com/Vallghall/book-list/pkg/models"
	"github.com/Vallghall/book-list/pkg/types"
	uuid "github.com/satori/go.uuid"
)

// CreateUser - user creation
func (s *DB) CreateUser(c *types.UserCreation) (*types.User, error) {
	u := &models.User{
		UserName:  c.UserName,
		FirstName: c.FirstName.NullString,
		LastName:  c.LastName.NullString,
		Email:     c.Email,
	}

	err := s.db.Create(u).Error

	return u.ToTarget(), err
}

// GetUser - get user by id
func (s *DB) GetUser(userID uuid.UUID) (*types.User, error) {
	u := &models.User{ID: userID}
	err := s.db.First(u).Error

	if err != nil {
		return nil, err
	}

	return u.ToTarget(), nil
}
