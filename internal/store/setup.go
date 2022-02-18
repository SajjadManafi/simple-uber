package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

// NewPostgresStore creates a new PostgresStore
func NewPostgresStore(dsn string) *PostgresStore {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return &PostgresStore{
		db: db,
	}
}
