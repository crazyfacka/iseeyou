package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/crazyfacka/iseeyou/server/api"
	"github.com/crazyfacka/iseeyou/server/commons"
	"github.com/crazyfacka/iseeyou/server/interpreter"
)

func main() {
	lock := make(chan bool, 1)
	sigRcv := make(chan os.Signal, 1)
	signal.Notify(sigRcv, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// Routine to catch and handle signals
	go func(sigRcv chan os.Signal) {
		for sig := range sigRcv {
			commons.Debug("Signal caught: %s", sig.String())
			commons.Debug("Terminating application...")
			// TODO Clean all
			lock <- true
		}
	}(sigRcv)

	i := interpreter.GetInterpreter()
	api.StartAPI(i)

	<-lock
}
