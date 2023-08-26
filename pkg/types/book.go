package types

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

// BookCreation - represents information needed to create a book
type BookCreation struct {
	AuthorIDs   []uuid.UUID `json:"author_ids"`
	Title       string      `json:"title"`
	ReleaseYear null.Int    `json:"release_year"`
}

// Book - holds info about a book
type Book struct {
	ID          uuid.UUID     `json:"id"`
	Title       string        `json:"title"`
	ReleaseYear null.Int      `json:"release_year"`
	Authors     []*AuthorLink `json:"authors,omitempty"`
}
