package handler

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/markraiter/chat/models"
)

func (h *Handler) isUserBlocked(userID, blockedUserID int) bool {
	return h.services.IsUserBlocked(userID, blockedUserID)
}

func (h *Handler) broadcastMessage(message models.Message) {
	for client := range h.clients {
		err := client.WriteJSON(message)
		if err != nil {
			log.Printf("error writing message: %s", err.Error())
			client.Close()
			delete(h.clients, client)
		}
	}
}

func (h *Handler) websocketHandler(c echo.Context) error {
	conn, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return fmt.Errorf("error upgrading to WebSocket: %w", err)
	}
	defer conn.Close()

	h.clients[conn] = true

	for {
		var message models.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("error reading message: %s", err.Error())
			delete(h.clients, conn)
			break
		}

		if h.isUserBlocked(message.UserID, message.BlockedUserID) {
			continue
		}

		h.broadcast <- message
		h.broadcastMessage(message)
	}

	return nil
}
