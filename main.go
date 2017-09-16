package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var configuration Config
var loginServer LoginServer

func main() {

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Info("OpenLou has been started!")
	configuration.Load("settings.json")
	loadedConstructions.LoadAllConstructions()
	loginServer.StartAndListen(1234)

}
