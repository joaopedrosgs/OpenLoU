package main

import (
	"OpenLoU/configuration"
	"OpenLoU/hermes"
	"OpenLoU/loginserver"
	"OpenLoU/mapserver"
	log "github.com/sirupsen/logrus"
	"os"
)

var config configuration.Config

func main() {

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Info("OpenLou has been started!")

	config.Load("settings.json")

	LoginServer, err := loginserver.CreateAndConnect(&config)
	if err != nil {
		panic("Login server could not be started")
	}

	MapServer, err := mapserver.CreateAndConnect(&config)
	if err != nil {
		panic("Map server could not be started")
	}

	Hermes := hermes.Create(&MapServer.In, &LoginServer.In, &LoginServer.In)
	Hermes.StartListening()
}
