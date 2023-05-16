package config

import (
	"log"

	"github.com/markraiter/chat/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB  *gorm.DB
	Err error
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) ConnectToDB() {
	dsn := "root:example@tcp(db:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	db.DB, db.Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if db.Err != nil {
		log.Fatalf("error connecting to database: %s/n", db.Err.Error())
	}
}

func (db *Database) MakeMigrations() {
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Message{})
}
