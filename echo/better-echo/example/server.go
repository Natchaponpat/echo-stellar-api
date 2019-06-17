package main

import (
	"github.com/Natchaponpat/echo-stellar-api/echo/better-echo/example/handle/user"
	"github.com/Natchaponpat/echo-stellar-api/echo/better-echo/example/storage"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userStorage := storage.NewUserStorage()
	userHandler, err := user.New(userStorage)
	if err != nil {
		e.Logger.Errorf("cannot init user handler: %v", err)
		return
	}
	userHandler.Handle(e)

	e.Logger.Info("start server")
	e.Logger.Fatal(e.Start(":8000"))
}
