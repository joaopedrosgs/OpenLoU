package main

import (
	"github.com/joaopedrosgs/OpenLoU/loginserver"
	log "github.com/sirupsen/logrus"
	"os"
)

var configuration Config

func main() {

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Info("OpenLou has been started!")
	configuration.Load("settings.json")
	LoginServer := loginserver.New(1000)
	LoginServer.StartListening()

}
