package postgres

import (
	"fmt"

	"github.com/Vallghall/book-list/pkg/models"
	"github.com/Vallghall/book-list/pkg/types"
	uuid "github.com/satori/go.uuid"
)

// CreateUser - user creation
func (s *DB) CreateUser(c *types.UserCreation) (*types.User, error) {
	u := &models.User{
		UserName:     c.UserName,
		FirstName:    c.FirstName.NullString,
		LastName:     c.LastName.NullString,
		Email:        c.Email,
		PasswordHash: c.Password,
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

// FindUserByCredentials - fetches user by their username and password
func (s *DB) FindUserByCredentials(name, pw string) (*types.User, error) {
	fmt.Println(name, pw)
	user := new(models.User)
	err := s.db.Where(
		&models.User{
			UserName:     name,
			PasswordHash: pw,
		}).First(user).Error

	if err != nil {
		return nil, err
	}

	return user.ToTarget(), nil
}

// FindUserByUsername - fetches user by his username
func (s *DB) FindUserByUsername(username string) (*types.User, error) {
	user := &models.User{
		UserName: username,
	}
	err := s.db.Where(user).First(user).Error
	if err != nil {
		return nil, err
	}

	return user.ToTarget(), nil
}
