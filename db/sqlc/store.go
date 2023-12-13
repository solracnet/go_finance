package db

import (
	"database/sql"
)

type Store interface {
	Querier
}

type SqlStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) *SqlStore {
	return &SqlStore{
		db:      db,
		Queries: New(db),
	}
}
