package main

var configuration Config
var loginServer LoginServer

func main() {
	configuration.Load("settings.json")
	loadedConstructions.LoadAllConstructions()
	loginServer.StartAndListen(1234)

}
