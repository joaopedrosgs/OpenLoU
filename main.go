package main

import (
	"OpenLoU/hermes"
	"OpenLoU/loginserver"
	"OpenLoU/mapserver"
	"OpenLoU/session"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime/pprof"
)

func main() {
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Info("OpenLou has been started!")

	Sessions := session.New()
	Hermes := hermes.Create(Sessions)

	LoginServer, err := loginserver.New(Sessions)
	if err != nil {
		panic("Email server could not be started")
	}

	MapServer, err := mapserver.New()
	if err != nil {
		panic("Map server could not be started")
	}

	Hermes.RegisterWorker(MapServer)

	go MapServer.StartListening()
	go LoginServer.StartListening()

	Hermes.StartListening()

}
