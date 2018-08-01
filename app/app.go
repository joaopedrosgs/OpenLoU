package app

import (
	log "github.com/sirupsen/logrus"

	"github.com/joaopedrosgs/OpenLoU/accountserver"
	"github.com/joaopedrosgs/OpenLoU/authserver"
	"github.com/joaopedrosgs/OpenLoU/cityserver"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/hub"
	"github.com/joaopedrosgs/OpenLoU/mapserver"
	"github.com/joaopedrosgs/OpenLoU/modules"
	"github.com/joaopedrosgs/OpenLoU/session"
	"os"
)

func Run() {
	context := log.WithField("Entity", "OpenLoU")
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	context.Info("OpenLoU is starting...")
	configuration.Load()
	MapServer := mapserver.New()
	CityServer := cityserver.New()
	AccountServer := accountserver.New()
	Hub := hub.New()
	database.Open()
	defer database.Close()

	modules.RegisterAllTroops()
	modules.RegisterAllConstructions()
	session.NewSessionInMemory()

	AuthServer, err := authserver.New()
	if err != nil {
		context.Error(err.Error())
	}
	err = Hub.RegisterServer(CityServer)
	if err != nil {
		context.Error(err.Error())
	}
	err = Hub.RegisterServer(MapServer)
	if err != nil {
		context.Error(err.Error())
	}
	err = Hub.RegisterServer(AccountServer)
	if err != nil {
		context.Error(err.Error())
	}

	go MapServer.StartListening()
	go CityServer.StartListening()
	go AccountServer.StartListening()
	go AuthServer.StartListening(":8000")

	Hub.StartListening(":8080")

}
