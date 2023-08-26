package models

import (
	"database/sql"

	"github.com/Vallghall/book-list/pkg/types"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

// Book - DAO for books
type Book struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string
	ReleaseYear sql.NullInt64
	Authors     []*Author `gorm:"many2many:book_authorship"`
	CreatedByID uuid.UUID `gorm:"type:uuid"`
	CreatedBy   *User
}

// ToTarget - transforms *models.Book to *types.Book
func (b *Book) ToTarget() *types.Book {
	authors := make([]*types.AuthorLink, len(b.Authors))
	for i, a := range b.Authors {
		authors[i] = a.ToTargetLink()
	}

	return &types.Book{
		ID:          b.ID,
		Title:       b.Title,
		ReleaseYear: null.IntFrom(b.ReleaseYear.Int64),
		Authors:     authors,
	}
}
