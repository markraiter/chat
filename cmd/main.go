package main

import (
	"log"

	"github.com/markraiter/chat/internal/handler"
	"github.com/markraiter/chat/internal/storage"
	"github.com/markraiter/chat/internal/storage/mysql"
	"github.com/markraiter/chat/models"
	"github.com/spf13/viper"
)

func main() {
	// Initialization of .yml config file
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	// Initialization of the MySQL database
	db, err := mysql.NewMySQL(mysql.Config{
		Driver:     viper.GetString("db.driver"),
		Username:   viper.GetString("db.username"),
		Password:   viper.GetString("db.password"),
		Connection: viper.GetString("db.connection"),
		Host:       viper.GetString("db.host"),
		Port:       viper.GetString("db.port"),
		DBName:     viper.GetString("db.dbname"),
	})

	if err != nil {
		log.Fatalf("Error initializing database: %s", err.Error())
	}

	// Making migrations using GORM
	db.AutoMigrate(&models.User{})

	// Initialization of the storage
	s := storage.NewStorage(db)
	// Initialization of the handler
	h := handler.NewHandler(s)

	// Initialization of the router and routes
	e := h.InitRoutes()

	// Starting the server
	e.Logger.Fatal(e.Start(viper.GetString("port")))
}

// initConfig initializes .yml configuration file for the chat application
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
