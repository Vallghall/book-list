package postgres

import (
	"github.com/Vallghall/book-list/pkg/store"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) store.Store {
	return &DB{db}
}
