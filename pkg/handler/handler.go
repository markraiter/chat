package handler

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/models"
	"github.com/markraiter/chat/pkg/service"
)

type Handler struct {
	services  *service.Service
	upgrader  *websocket.Upgrader
	clients   map[*websocket.Conn]bool
	broadcast chan models.Message
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services:  services,
		upgrader:  &websocket.Upgrader{},
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan models.Message),
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()

	e.GET("/ws", h.websocketHandler)

	e.POST("/login", h.login)
	e.POST("/register", h.register)

	api := e.Group("/api")
	api.Use(h.JWTMiddleware)
	friendList := api.Group("/friendlist")
	blacklist := api.Group("/blacklist")

	friendList.POST("/", h.addToFriends)
	friendList.DELETE("/:id", h.deleteFriend)

	blacklist.POST("/", h.addToBlacklist)
	blacklist.DELETE("/:id", h.deleteFromBlacklist)

	return e
}
