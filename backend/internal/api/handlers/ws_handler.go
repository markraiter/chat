package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/markraiter/chat/internal/models"
	"github.com/markraiter/chat/internal/util"
)

type WSHandler struct {
	hub *models.Hub
}

func NewWSHandler(h *models.Hub) *WSHandler {
	return &WSHandler{hub: h}
}

type CreateRoomReq struct {
	ID   string `json:"id" example:"1"`
	Name string `json:"name" example:"Room_1"`
}

// @Summary CreateRoom
// @Tags Websocket
// @Description create room
// @ID create-room
// @Accept  json
// @Produce  json
// @Param input body CreateRoomReq true "room info"
// @Success 201 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 406 {object} util.Response
// @Router /ws/create-room [post].
func (h *WSHandler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.Response{Message: err.Error()})

		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusNotAcceptable, util.Response{Message: err.Error()})

		return
	}

	h.hub.Rooms[req.ID] = &models.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*models.Client),
	}

	c.JSON(http.StatusCreated, util.Response{Message: fmt.Sprintf("room %s created", req.Name)})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// origin := r.Header.Get("Origin")
		// return origin == "http://localhost:3000"

		return true
	},
}

// @Summary JoinRoom
// @Tags Websocket
// @Description join room
// @ID join-room
// @Produce  json
// @Param room_id path string true "room_id"
// @Param user_id query string true "user_id"
// @Param username query string true "username"
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Router /ws/join-room [get].
func (h *WSHandler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Response{Message: err.Error()})

		return
	}

	roomID := c.Param("room_id")
	clientID := c.Query("user_id")
	username := c.Query("username")

	cl := &models.Client{
		Conn:     conn,
		Message:  make(chan *models.Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	m := &models.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// @Summary GetRooms
// @Tags Websocket
// @Description get all rooms
// @ID get-rooms
// @Produce  json
// @Success 200 {object} util.Response
// @Router /ws/get-rooms [get].
func (h *WSHandler) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// @Summary GetClients
// @Tags Websocket
// @Description get all clients in the room
// @ID get-clients
// @Produce  json
// @Param room_id path string true "room_id"
// @Success 200 {object} util.Response
// @Router /ws/get-clients [get].
func (h *WSHandler) GetClients(c *gin.Context) {
	var clients []ClientRes
	roomID := c.Param("room_id")

	if _, ok := h.hub.Rooms[roomID]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)

		return
	}

	for _, c := range h.hub.Rooms[roomID].Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}
