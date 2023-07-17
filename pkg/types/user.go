package types

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

// UserCreation - data for user creation
type UserCreation struct {
	Nickname  null.String `json:"nickname"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
}

// User - represents info about users
type User struct {
	ID         uuid.UUID   `json:"id" db:"id"`
	Nickname   null.String `json:"nickname" db:"nickname"`
	FirstName  string      `json:"first_name" db:"first_name"`
	LastName   string      `json:"last_name" db:"last_name"`
	Email      string      `json:"email" db:"email"`
	CreatedAt  time.Time   `json:"created_at" db:"created_at"`
	ModifiedAt time.Time   `json:"modified_at" db:"modified_at"`
}
