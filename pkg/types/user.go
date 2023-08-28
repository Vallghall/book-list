package types

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

// UserCreation - data for user creation
type UserCreation struct {
	UserName  string      `json:"username"`
	FirstName null.String `json:"first_name"`
	LastName  null.String `json:"last_name"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
}

// User - represents info about users
type User struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"username"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
}
