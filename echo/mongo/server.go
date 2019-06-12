package main

import (
	"github.com/Natchaponpat/echo-stellar-api/echo/mongo/handle/user"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := user.DBConfig{
		Host:      "localhost:27017",
		Username:  "admin",
		Pass:      "password",
		DB:        "test",
		Colletion: "users",
	}
	userHandler, err := user.New(config)
	if err != nil {
		e.Logger.Errorf("cannot init user handler: %v", err)
		return
	}
	userHandler.Handle(e)

	e.Logger.Info("start server")
	e.Logger.Fatal(e.Start(":8000"))
}
