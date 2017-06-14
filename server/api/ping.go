package api

import (
	"net/http"

	"github.com/crazyfacka/iseeyou/server/commons"
	"github.com/labstack/echo"
)

func (api *API) ping(c echo.Context) error {
	var msg []byte
	var err error

	msg, err = api.parseBody(c)
	commons.Debug("[PING] Received ping message: '%s'", string(msg))

	if err == nil {
		if api.i.StoreAlive(msg) {
			return c.String(http.StatusOK, commons.GetJSONMessage("pong"))
		}
	}

	return c.String(http.StatusBadRequest, commons.GetJSONMessage(err.Error()))
}
