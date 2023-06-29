package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()

	auth := e.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/login", h.login)
	}

	api := e.Group("/api")
	{
		friendList := api.Group("/friendlist")
		{
			friendList.POST("/", h.addToFriends)
			friendList.DELETE("/:id", h.deleteFriend)
		}

		blacklist := api.Group("/blacklist")
		{
			blacklist.POST("/", h.addToBlacklist)
			blacklist.DELETE("/:id", h.deleteFromBlacklist)
		}
	}

	return e
}
