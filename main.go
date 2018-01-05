package main

import (
	"github.com/joaopedrosgs/OpenLoU/hermes"
	"github.com/joaopedrosgs/OpenLoU/loginserver"
	"github.com/joaopedrosgs/OpenLoU/mapserver"
	"github.com/joaopedrosgs/OpenLoU/session"
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
