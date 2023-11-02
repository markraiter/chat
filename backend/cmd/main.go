package main

import (
	"fmt"
	"log"

	_ "github.com/markraiter/chat/docs"

	"github.com/markraiter/chat/internal/api"
	"github.com/markraiter/chat/internal/api/handlers"
	"github.com/markraiter/chat/internal/configs"
	"github.com/markraiter/chat/internal/models"
	"github.com/markraiter/chat/internal/service"
	"github.com/markraiter/chat/internal/storage/postgres"
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

	userRepository := postgres.NewRepository(dbConn.GetDB())
	userService := service.NewService(userRepository)
	userHandler := handlers.NewHandler(userService)

	hub := models.NewHub()
	wsHandler := handlers.NewWSHandler(hub)
	go hub.Run()

	api.InitRoutes(cfg, userHandler, wsHandler)

	if err := api.Start(cfg.Server); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
