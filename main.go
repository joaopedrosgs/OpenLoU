package main

import (
	"OpenLoU/hermes"
	"OpenLoU/loginserver"
	"OpenLoU/mapserver"
	"OpenLoU/session"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Info("OpenLou has been started!")

	Sessions := session.New()

	LoginServer, err := loginserver.New(Sessions)
	if err != nil {
		panic("Email server could not be started")
	}

	MapServer, err := mapserver.New()
	if err != nil {
		panic("Map server could not be started")
	}

	Hermes := hermes.Create(MapServer.GetEntryPoint(), Sessions)

	MapServer.SetEndPoint(Hermes.GetEntryPoint())
	go MapServer.StartListening()
	go LoginServer.StartListening()

	Hermes.StartListening()

}
