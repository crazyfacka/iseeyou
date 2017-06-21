package api

import (
	"net/http"

	"github.com/crazyfacka/iseeyou/server/commons"
	"github.com/labstack/echo"
)

func (api *API) getLatest(c echo.Context) error {
	if data, err := api.i.GetLatest(); err == nil {
		return c.String(http.StatusOK, data)
	}

	return c.String(http.StatusServiceUnavailable, commons.GetJSONMessage("Error retrieving information"))
}
