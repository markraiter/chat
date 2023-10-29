package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/markraiter/chat/internal/configs"
)

type Database struct {
	db *sql.DB
}

func NewDB(cfg configs.Postgres) (*Database, error) {
	db, err := sql.Open(cfg.Driver, cfg.ConnString)
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
