package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/internal/storage"
)

// Handler struct describes Handler entity
type Handler struct {
	storage *storage.Storage
}

// NewHandler function is a constructor function for Handler struct
func NewHandler(storage *storage.Storage) *Handler {
	return &Handler{storage: storage}
}

// InitRoutes function initialises new router and routes
func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()

	e.POST("/sign-up", h.signUp)
	e.POST("/sign-in", h.signIn)

	return e
}
