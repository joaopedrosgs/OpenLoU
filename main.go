package main

import (
	"OpenLoU/hermes"
	"OpenLoU/loginserver"
	"OpenLoU/mapserver"
	"OpenLoU/session"
)

func main() {

	println("OpenLou has been started!")

	Sessions := session.NewSessionInMemory()
	Hermes := hermes.Create(Sessions)

	LoginServer, err := loginserver.New(Sessions)
	if err != nil {
		println(err.Error())
	}

	MapServer, err := mapserver.New()
	if err != nil {
		println(err.Error())
	}

	Hermes.RegisterWorker(MapServer)

	go MapServer.StartListening()
	go LoginServer.StartListening()

	Hermes.StartListening()

}
