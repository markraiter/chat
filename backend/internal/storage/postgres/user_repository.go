package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/markraiter/chat/internal/models"
	"github.com/markraiter/chat/internal/util"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) *repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	var lastInsertID int
	query := "INSERT INTO users(username, password, email) VALUES ($1, $2, $3) returning id"
	if err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertID); err != nil {
		return nil, fmt.Errorf("user_repository CreateUser() error: %w", err)
	}

	user.ID = int64(lastInsertID)

	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	u := models.User{}

	query := "SELECT id, email, username, password FROM users WHERE email = $1"
	if err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrWrongCredentials
		}
		return nil, fmt.Errorf("user_repository GetUserByEmail() error: %w", err)
	}

	return &u, nil
}

func (r *repository) GetEmail(ctx context.Context, email string) string {
	u := models.User{}

	query := "SELECT id, email, username, password FROM users WHERE email ILIKE $1"
	if err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password); err != nil {
		return ""
	}

	return u.Email
}

func (r *repository) GetUsername(ctx context.Context, username string) string {
	u := models.User{}

	query := "SELECT id, email, username, password FROM users WHERE username = $1"
	if err := r.db.QueryRowContext(ctx, query, username).Scan(&u.ID, &u.Email, &u.Username, &u.Password); err != nil {
		return ""
	}

	return u.Username
}
