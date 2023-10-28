package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDB() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/chat-app-db?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("storage NewDB() error: %w", err)
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
