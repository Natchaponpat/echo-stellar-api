package main

import (
	"fmt"

	"github.com/Natchaponpat/echo-stellar-api/echo/mongo/handle/user"
	"github.com/Natchaponpat/echo-stellar-api/echo/mongo/storage"
	"github.com/globalsign/mgo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := storage.DBConfig{
		Host:      "localhost:27017",
		Username:  "admin",
		Pass:      "password",
		DB:        "test",
		Colletion: "users",
	}
	uri := fmt.Sprintf("%v:%v@%v", config.Username, config.Pass, config.Host)
	session, err := mgo.Dial(uri)
	if err != nil {
		e.Logger.Errorf("cannot connect mongo db: %v", err)
		return
	}
	userStorage := storage.NewUserStorage(session, config.DB, config.Colletion)

	userHandler, err := user.New(userStorage)
	if err != nil {
		e.Logger.Errorf("cannot init user handler: %v", err)
		return
	}
	userHandler.Handle(e)

	e.Logger.Info("start server")
	e.Logger.Fatal(e.Start(":8000"))
}
