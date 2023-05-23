package repository

import (
	"database/sql"
	"fmt"

	"github.com/markraiter/chat/models"
)

type AuthMySQL struct {
	db *sql.DB
}

func NewAuthMySQL(db *sql.DB) *AuthMySQL {
	return &AuthMySQL{
		db: db,
	}
}

func (r *AuthMySQL) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (avatar, username, password) values (?, ?, ?) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Avatar, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthMySQL) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = ? AND password = ?", usersTable)
	err := r.db.QueryRow(query, username, password).Scan(&user.ID, &user.Username, &user.Password)

	return user, err
}
