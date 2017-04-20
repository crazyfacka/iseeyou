package interpreter

import (
	"encoding/json"

	"github.com/crazyfacka/iseeyou/server/commons"
	"github.com/crazyfacka/iseeyou/server/handler"
)

// Interpreter core struct
type Interpreter struct {
	sqlh *handler.SQL
}

var interpreter *Interpreter

// StoreAlive stores an alive message
func (i *Interpreter) StoreAlive(msg []byte) bool {
	var ok bool
	var data map[string]interface{}

	if err := json.Unmarshal(msg, &data); err != nil {
		return false
	}

	var val float64
	if val, ok = data["start"].(float64); !ok {
		return false
	}

	return i.sqlh.StorePing(val)
}

// GetInterpreter intatiates this interpreter
func GetInterpreter() *Interpreter {
	interpreter := &Interpreter{
		sqlh: handler.GetSQLHandler(commons.GetConfiguration()),
	}

	return interpreter
}
