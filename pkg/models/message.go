package models

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	upgrader  = websocket.Upgrader{}
)

type Message struct {
	gorm.Model
	Username string
	Body     string
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) HandleWS(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	defer conn.Close()

	clients[conn] = true

	for {
		err := conn.ReadJSON(&m)
		if err != nil {
			delete(clients, conn)
			break
		}

		broadcast <- *m
	}

	return nil
}

func (m *Message) HandleMSG() {
	for {
		message := <-broadcast

		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
