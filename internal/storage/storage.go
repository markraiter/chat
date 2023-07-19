package storage

import (
	"github.com/markraiter/chat/internal/storage/mysql"
	"gorm.io/gorm"
)

// Storage struct contains all interfaces to work with the storage
type Storage struct {
	mysql.Auth
}

// NewStorage function is a constructor for Storage struct
func NewStorage(db *gorm.DB) *Storage {
	return &Storage{
		*mysql.NewAuth(db),
	}
}
