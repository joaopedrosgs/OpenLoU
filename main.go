package main

import (
	"LordOfUltima/config"
	"LordOfUltima/database"
)

Config config

func main() {
	config.Load("settings.json")
	constructions.RegisterAll()

	military.RegisterAll()
	MapServer.LoadAndStart()

}
