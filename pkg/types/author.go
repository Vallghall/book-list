package types

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

// AuthorCreation - represents info needed to create an author
type AuthorCreation struct {
	Name       string      `json:"name"`
	LastName   string      `json:"last_name"`
	FatherName null.String `json:"father_name"`
	Birthday   null.Time   `json:"birthday"`
}

// Author - author representation
type Author struct {
	ID         uuid.UUID   `json:"id"`
	Name       string      `json:"name"`
	LastName   string      `json:"last_name"`
	FatherName null.String `json:"father_name"`
	Birthday   null.Time   `json:"birthday"`
	CreatedBy  uuid.UUID   `json:"created_by"`
}

// AuthorLink - basic info about an author
type AuthorLink struct {
	ID         uuid.UUID   `json:"id"`
	Name       string      `json:"name"`
	LastName   string      `json:"last_name"`
	FatherName null.String `json:"father_name"`
}
