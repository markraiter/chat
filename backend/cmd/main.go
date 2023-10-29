package main

import (
	"fmt"
	"log"

	_ "github.com/markraiter/chat/docs"

	"github.com/markraiter/chat/internal/configs"
	"github.com/markraiter/chat/internal/router"
	"github.com/markraiter/chat/internal/storage/postgres"
	"github.com/markraiter/chat/internal/user"
	"github.com/markraiter/chat/internal/websocket"
)

//	@title			CHAT APP
//	@version		1.0
//	@description	Docs for chat app backend API
//	@contact.name	Mark Raiter
//	@contact.email	raitermark@proton.me
//  host  			localhost:9000
//	@BasePath		/

func main() {
	cfg, err := configs.InitConfig()
	if err != nil {
		log.Fatalf("InitConfig error: %s\n", err.Error())

		return
	}

	fmt.Printf("%+v\n", cfg)

	dbConn, err := postgres.NewDB(cfg.DB)
	if err != nil {
		log.Fatalf("could not initialize database connection: %s\n", err.Error())
	}

	userRepository := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	hub := websocket.NewHub()
	wsHandler := websocket.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start(cfg.Server.AppAddress)
}
