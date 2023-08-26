package models

import (
	"database/sql"

	"github.com/Vallghall/book-list/pkg/types"
	uuid "github.com/satori/go.uuid"
)

// User - DAO for users
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserName  string    `gorm:"unique"`
	FirstName sql.NullString
	LastName  sql.NullString
	Email     string `gorm:"check:(email like '%@%.%')"`
}

// ToTarget - transforms *models.User to *types.User
func (u *User) ToTarget() *types.User {
	return &types.User{
		ID:        u.ID,
		UserName:  u.UserName,
		FirstName: u.FirstName.String,
		LastName:  u.LastName.String,
		Email:     u.Email,
	}
}
