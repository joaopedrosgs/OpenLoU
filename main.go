package main

import (
	"github.com/joaopedrosgs/OpenLoU/accountserver"
	"github.com/joaopedrosgs/OpenLoU/cityserver"
	"github.com/joaopedrosgs/OpenLoU/entities"
	"github.com/joaopedrosgs/OpenLoU/hub"
	"github.com/joaopedrosgs/OpenLoU/mapserver"
	"github.com/joaopedrosgs/OpenLoU/session"

	log "github.com/sirupsen/logrus"
	"os"

	"github.com/joaopedrosgs/OpenLoU/database"
)

var context = log.WithField("Entity", "OpenLoU")

func main() {

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	context.Info("OpenLoU is starting...")
	database.InitDB()
	entities.RegisterAllTroops()
	entities.RegisterAllConstructions()

	Sessions := session.NewSessionInMemory()
	Hermes := hub.Create(Sessions)

	LoginServer, err := accountserver.New(Sessions)
	if err != nil {
		context.Error(err.Error())
	}

	MapServer := mapserver.New()
	CityServer := cityserver.New()

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

	Hermes.StartListening(":8080")

}
