package main

import (
	"fmt"
	"net/http"

	"github.com/markraiter/chat/pkg/websocket"
)

func main() {
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}

func setupRoutes() {
	pool := websocket.NewPool()

	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
