package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/crazyfacka/iseeyou/server/commons"
	"github.com/crazyfacka/iseeyou/server/interpreter"
	"github.com/labstack/echo"
)

// API core struct
type API struct {
	port int64
	i    *interpreter.Interpreter
}

func (api *API) parseBody(c echo.Context) ([]byte, error) {
	var msg []byte
	var err error

	if msg, err = ioutil.ReadAll(c.Request().Body); err == nil {
		commons.Debug("[PING] Received ping message: '%s'", string(msg))
		return msg, nil
	}

	return nil, errors.New("error processing request")
}

func (api *API) bananas() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "These are not the droids you are looking for...")
	})

	e.GET("/latest", api.getLatest)
	e.PUT("/motion", api.motion)
	e.POST("/ping", api.ping)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(int(api.port))))
}

// StartAPI starts this API
func StartAPI(i *interpreter.Interpreter) *API {
	cfg := commons.GetConfiguration()

	api := &API{
		port: cfg.API.Port,
		i:    i,
	}

	go api.bananas()
	return api
}
