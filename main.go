package main

import (
	"github.com/joaopedrosgs/OpenLoU/cityserver"
	"github.com/joaopedrosgs/OpenLoU/hermes"
	"github.com/joaopedrosgs/OpenLoU/loginserver"
	"github.com/joaopedrosgs/OpenLoU/mapserver"
	"github.com/joaopedrosgs/OpenLoU/session"
	log "github.com/sirupsen/logrus"
	"os"
)

var context = log.WithField("Entity", "OpenLoU")

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	context.Info("OpenLoU is starting...")

	Sessions := session.NewSessionInMemory()
	Hermes := hermes.Create(Sessions)

	LoginServer, err := loginserver.New(Sessions)
	if err != nil {
		context.Error(err.Error())
	}

	MapServer, err := mapserver.New()
	if err != nil {
		context.Error(err.Error())
	}

	CityServer, err := cityserver.New()
	if err != nil {
		context.Error(err.Error())
	}
	err = Hermes.RegisterWorker(CityServer)
	if err != nil {
		context.Error(err.Error())
	}
	err = Hermes.RegisterWorker(MapServer)
	if err != nil {
		context.Error(err.Error())
	}

	go MapServer.StartListening()
	go LoginServer.StartListening(":8000")

	Hermes.StartListening()

}
