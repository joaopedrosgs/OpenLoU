package main

import (
	"LordOfUltima/configuration"
	"LordOfUltima/constructions"
	"LordOfUltima/database"
	"LordOfUltima/military"
)

var MapServer mapServer

func StartUp() {
	configuration.Load("settings.json")
	database.Open()
	constructions.RegisterAll()
	military.RegisterAll()
	MapServer.LoadAndStart()

}

func main() {
	StartUp()

}
