package api

import (
	"net/http"

	"github.com/crazyfacka/iseeyou/server/commons"
	"github.com/labstack/echo"
)

func (api *API) getLatest(c echo.Context) error {
	return c.String(http.StatusServiceUnavailable, commons.GetJSONMessage("Error retrieving information"))
}
