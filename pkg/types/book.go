package types

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

// BookCreation - represents information needed to create a book
type BookCreation struct {
	Title       string   `json:"title"`
	ReleaseYear null.Int `json:"release_year"`
}

// Book - holds info about a book
type Book struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	ReleaseYear null.Int  `json:"release_year" db:"release_year"`
	*AuthorLink
}
