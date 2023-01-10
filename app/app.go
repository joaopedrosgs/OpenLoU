package app

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"

	"openlou/accountserver"
	"openlou/cityserver"
	"openlou/hub"
	"openlou/mapserver"
	"openlou/modules"
	"openlou/session"
)

func Run() {
	context := log.WithField("Entity", "OpenLoU")
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	context.Info("OpenLoU is starting...")
	MapServer := mapserver.New()
	CityServer := cityserver.New()
	AccountServer := accountserver.New()

	Hub, err := hub.New()
	if err != nil {
		context.Fatal(err.Error())
		return
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
	if viper.GetBool("server.debug") {
		AccountServer.CreateAdminAccountAction()
	}
	Hub.Start()
}
