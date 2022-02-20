package store

import (
	"database/sql"
	"fmt"

	"github.com/SajjadManafi/simple-uber/internal/config"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

// NewPostgresStore creates a new PostgresStore
func NewPostgresStore(config config.Config) (*PostgresStore, error) {
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return &PostgresStore{
		db: db,
	}, nil
}
