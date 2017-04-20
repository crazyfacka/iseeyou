package main

import (
	"io/ioutil"
	"net/http"

	"github.com/crazyfacka/iseeyou/server/commons"
	"github.com/crazyfacka/iseeyou/server/interpreter"

	"github.com/labstack/echo"
)

func main() {
	i := interpreter.GetInterpreter()
	e := echo.New()

	e.POST("/ping", func(c echo.Context) error {
		var msg []byte
		var err error

		if msg, err = ioutil.ReadAll(c.Request().Body); err == nil {
			commons.Debug("[PING] Received ping message: '%s'", string(msg))
			if i.StoreAlive(msg) {
				return c.String(http.StatusOK, commons.GetJSONMessage("pong"))
			}
		}

		return c.String(http.StatusBadRequest, commons.GetJSONMessage("error processing request"))
	})

	e.PUT("/motion", func(c echo.Context) error {
		var msg []byte
		var err error

		if msg, err = ioutil.ReadAll(c.Request().Body); err == nil {
			commons.Debug("[MOTION] Received motion message: '%s'", string(msg))
			if i.StoreMotion(msg) {
				return c.String(http.StatusOK, commons.GetJSONMessage("ok"))
			}
		}

		return c.String(http.StatusBadRequest, commons.GetJSONMessage("error processing request"))
	})

	e.Logger.Fatal(e.Start(":9000"))
}
