package repository

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	usersTable    = "users"
	messagesTable = "messages"
)

type Config struct {
	Driver     string
	Username   string
	Password   string
	Connection string
	Host       string
	Port       string
	DBName     string
}

func NewMySQLDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf("%s:%s@%s(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Connection, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		log.Fatalf("error connecting to database: %s\n", err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
