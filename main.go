package main

import (
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/loginserver"
	log "github.com/sirupsen/logrus"
	"os"
)

var config configuration.Config

func main() {

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Info("OpenLou has been started!")
	config.Load("settings.json")
	LoginServer := loginserver.New(true, &config)
	LoginServer.StartListening()

}
