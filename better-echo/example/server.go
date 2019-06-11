package main

import (
	"github.com/Natchaponpat/echo-stellar-api/better-echo/example/handle/user"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))

	userHandler, err := user.New()
	if err != nil {
		e.Logger.Errorf("cannot init user handler: %v", err)
		return
	}
	userHandler.Handle(e)

	e.Logger.Info("start server")
	e.Logger.Fatal(e.Start(":8000"))
}