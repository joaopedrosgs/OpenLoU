package main

import (
	"LordOfUltima/config"
	"LordOfUltima/constructions"
	"LordOfUltima/database"
	"LordOfUltima/military"
)

var MapServer mapServer

func StartUp() {
	config.Load()
	database.Open()
	constructions.RegisterAll()
	military.RegisterAll()
	MapServer.LoadAndStart()

}

func main() {
	StartUp()

}
