package types

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4"
)

// AuthorCreation - represents info needed to create an author
type AuthorCreation struct {
	Name       string      `json:"name"`
	LastName   string      `json:"last_name"`
	FatherName null.String `json:"father_name"`
	Birthday   time.Time   `json:"birthday"`
}

// Author - author representation
type Author struct {
	ID         uuid.UUID   `json:"id" db:"id"`
	Name       string      `json:"name" db:"name"`
	LastName   string      `json:"last_name" db:"last_name"`
	FatherName null.String `json:"father_name" db:"father_name"`
	Birthday   time.Time   `json:"birthday" db:"birthday"`
	CreatedAt  time.Time   `json:"created_at" db:"created_at"`
	CreatedBy  uuid.UUID   `json:"created_by" db:"created_by"`
	ModifiedAt time.Time   `json:"modified_at" db:"modified_at"`
	ModifiedBy uuid.UUID   `json:"modified_by" db:"modified_by"`
}

// AuthorLink - basic info about an author
type AuthorLink struct {
	ID         uuid.UUID   `json:"id" db:"id"`
	Name       string      `json:"name" db:"name"`
	LastName   string      `json:"last_name" db:"last_name"`
	FatherName null.String `json:"father_name" db:"father_name"`
}
