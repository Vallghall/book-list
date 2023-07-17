package postgres

import (
	"github.com/Vallghall/book-list/pkg/types"
	uuid "github.com/satori/go.uuid"
)

// CreateUser - user creation
func (s *DB) CreateUser(c *types.UserCreation) (*types.User, error) {
	u := new(types.User)
	err := s.db.Get(u,
		"select * from k_user.create_user(($1,$2,$3,$4)::k_user._user);",
		c.Nickname, c.FirstName, c.LastName, c.Email,
	)

	return u, err
}

// GetUser - get user by id
func (s *DB) GetUser(userID uuid.UUID) (*types.User, error) {
	u := new(types.User)
	err := s.db.Get(u, "select * from k_user.get_user_by_id($1::uuid);", userID)

	return u, err
}
