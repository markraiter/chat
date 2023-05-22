package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markraiter/chat/pkg/models"
	"github.com/markraiter/chat/pkg/repo"
)

var (
	msg = models.NewMessage()
	db  = repo.NewDatabase()
)

func Start() {
	//setup
	db.ConnectToDB()
	db.MakeMigrations()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// websocket endpoint
	e.GET("/ws", msg.HandleWS)
	go msg.HandleMSG()

	// db endpoints
	e.POST("/register", db.Register)
	e.POST("/login", db.Login)
	e.GET("/users", db.GetUsers)
	e.GET("/users/:id", db.GetUserByID)

	e.Logger.Fatal(e.Start(":1323"))
}
