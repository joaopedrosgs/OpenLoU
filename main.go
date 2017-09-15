package main

import ()

var configuration Config
var loginServer LoginServer

func main() {
	configuration.Load("settings.json")
	loadedConstructions.LoadAllConstructions()
	LoginServer.StartAndListen(1234)

}
