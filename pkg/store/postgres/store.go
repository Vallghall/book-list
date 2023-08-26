package postgres

import (
	"github.com/Vallghall/book-list/pkg/store"
	"gorm.io/gorm"
)

// interface guard
var _ store.Store = (*DB)(nil)

type DB struct {
	db *gorm.DB
}

func New(db *gorm.DB) store.Store {
	return &DB{db}
}
