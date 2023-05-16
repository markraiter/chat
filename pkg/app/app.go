package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markraiter/chat/config"
	"github.com/markraiter/chat/pkg/models"
)

var (
	msg = models.NewMessage()
	usr = models.NewUser()
	db  = config.NewDatabase()
)

func Start() {
	//setup
	db.ConnectToDB()
	// db.MakeMigrations()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// websocket endpoint
	e.GET("/ws", msg.HandleWS)
	go msg.HandleMSG()

	// users endpoints
	e.GET("/users", usr.GetUsers(db.DB))
	e.GET("/users/:id", usr.GetUserByID(db.DB))
	e.POST("/register", usr.Register(db.DB))
	e.POST("/login", usr.Login(db.DB))

	e.Logger.Fatal(e.Start(":1323"))
}
