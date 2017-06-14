package api

import (
	"net/http"

	"github.com/crazyfacka/iseeyou/server/commons"
	"github.com/labstack/echo"
)

func (api *API) motion(c echo.Context) error {
	var msg []byte
	var err error

	msg, err = api.parseBody(c)
	commons.Debug("[MOTION] Received motion message: '%s'", string(msg))

	if err == nil {
		if api.i.StoreMotion(msg) {
			return c.String(http.StatusOK, commons.GetJSONMessage("ok"))
		}
	}

	return c.String(http.StatusBadRequest, commons.GetJSONMessage("error processing request"))
}
