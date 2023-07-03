package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func NewMySQLDB(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@%s(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Connection, cfg.Host, cfg.Port, cfg.DBName)), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database: %s\n", err.Error())
	}

	return db, nil
}
