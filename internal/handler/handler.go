package handler

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/internal/storage"
	"github.com/markraiter/chat/models"
)

// Handler struct describes Handler entity
type Handler struct {
	storage   *storage.Storage
	upgrader  *websocket.Upgrader
	clients   map[*websocket.Conn]bool
	broadcast chan models.Message
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

	apiGroup := e.Group("/api", h.JWTMiddleware)
	homePageGroup := apiGroup.Group("/home")
	// friendsGroup := apiGroup.Group("/friends")
	// blacklistGroup := apiGroup.Group("/blacklist")

	homePageGroup.PATCH("/:id", h.updateInfo) // updates users' info such as username or password
	homePageGroup.GET("", h.handleWS)

	// friendsGroup.POST("", h.addFriend)                     // adds friend
	// friendsGroup.GET("", h.showFriends)                    // gets all friends of current user
	// friendsGroup.DELETE("/:username", h.removeFromFriends) // deletes friend of current user

	// blacklistGroup.POST("", h.addToBalcklist)               // adds user to blacklist
	// blacklistGroup.GET("", h.getBlocked)                    // gets all blocked users of current user
	// blacklistGroup.DELETE("/username", h.removeFromBlocked) // removes user from blacklist

	return e
}
