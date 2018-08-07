package app

import (
	log "github.com/sirupsen/logrus"

	"os"

	"github.com/joaopedrosgs/OpenLoU/accountserver"
	"github.com/joaopedrosgs/OpenLoU/authserver"
	"github.com/joaopedrosgs/OpenLoU/cityserver"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/hub"
	"github.com/joaopedrosgs/OpenLoU/mapserver"
	"github.com/joaopedrosgs/OpenLoU/modules"
	"github.com/joaopedrosgs/OpenLoU/session"
)

func Run(port string) {
	context := log.WithField("Entity", "OpenLoU")
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	context.Info("OpenLoU is starting...")
	configuration.Load()
	MapServer := mapserver.New()
	CityServer := cityserver.New()
	AccountServer := accountserver.New()
	AuthServer := authserver.New()

	Hub, err := hub.New()
	if err != nil {
		context.Error(err.Error())
	}
	modules.RegisterAllTroops()
	modules.RegisterAllConstructions()
	session.NewSessionInMemory()

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
	err = Hub.RegisterServer(AuthServer)
	if err != nil {
		context.Error(err.Error())
	}
	Hub.Start(port)
}
