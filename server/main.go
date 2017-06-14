package main

import (
	"github.com/crazyfacka/iseeyou/server/api"
	"github.com/crazyfacka/iseeyou/server/interpreter"
)

func main() {
	i := interpreter.GetInterpreter()
	api.StartAPI(i)
}
