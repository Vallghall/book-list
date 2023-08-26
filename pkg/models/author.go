package models

import (
	"database/sql"

	"github.com/Vallghall/book-list/pkg/types"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

// Author - DAO for authors
type Author struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string
	LastName    string
	FatherName  sql.NullString
	Birthday    sql.NullTime
	CreatedByID uuid.UUID `gorm:"type:uuid"`
	CreatedBy   *User
}

// ToTarget - transforms *models.Author into *types.Author
func (a *Author) ToTarget() *types.Author {
	return &types.Author{
		ID:         a.ID,
		Name:       a.Name,
		LastName:   a.LastName,
		FatherName: null.StringFrom(a.FatherName.String),
		Birthday:   null.TimeFrom(a.Birthday.Time),
		CreatedBy:  a.CreatedByID,
	}
}

// ToTargetLink - transforms *models.Author into *types.AuthorLink,
// a shorter version of *types.Author
func (a *Author) ToTargetLink() *types.AuthorLink {
	return &types.AuthorLink{
		ID:         a.ID,
		Name:       a.Name,
		LastName:   a.LastName,
		FatherName: null.StringFrom(a.FatherName.String),
	}
}
