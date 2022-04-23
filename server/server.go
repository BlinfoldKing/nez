package server

import (
	"nez/handler"

	"github.com/labstack/echo/v4"
)

func Init() error {
	e := echo.New()

	e.GET("ping", handler.Handler(handler.Ping))
	e.GET("ping/error", handler.Handler(handler.PingErr))

	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
