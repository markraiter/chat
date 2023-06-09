package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/markraiter/chat/cmd/server"
	"github.com/markraiter/chat/pkg/handler"
	"github.com/markraiter/chat/pkg/repository"
	"github.com/markraiter/chat/pkg/service"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s\n", err.Error())
	}

	db, err := repository.NewMySQLDB(repository.Config{
		Driver:     viper.GetString("db.driver"),
		Username:   viper.GetString("db.username"),
		Password:   viper.GetString("db.password"),
		Connection: viper.GetString("db.connection"),
		Host:       viper.GetString("db.host"),
		Port:       viper.GetString("db.port"),
		DBName:     viper.GetString("db.dbname"),
	})

	if err != nil {
		log.Fatalf("error initializing database: %s\n", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s\n", err.Error())
		}
	}()

	log.Print("Chat Started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Chat Shutting Down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured while server shutting down: %s\n", err.Error())
	}
}
