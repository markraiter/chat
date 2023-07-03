package main

import (
	"log"

	"github.com/markraiter/chat/models"
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

	db.AutoMigrate(&models.User{}, &models.Message{}, &models.Friendship{}, &models.Blacklist{})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes().Logger.Fatal(handlers.InitRoutes().Start(viper.GetString("port")))
}
