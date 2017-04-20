package main

import (
	"net/http"

	"github.com/crazyfacka/iseeyou/server/commons"
	"github.com/crazyfacka/iseeyou/server/handler"

	"github.com/labstack/echo"
)

func main() {
	sqlh := handler.GetSQLHandler(commons.GetConfiguration())

	e := echo.New()

	e.POST("/ping", func(c echo.Context) error {
		commons.Debug("[PING] Received ping message: '%s'", c.QueryParams().Encode())
		sqlh.StorePing()
		return c.String(http.StatusOK, "pong")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
