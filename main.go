package main

import (
	"github.com/joaopedrosgs/OpenLoU/accountserver"
	"github.com/joaopedrosgs/OpenLoU/cityserver"
	"github.com/joaopedrosgs/OpenLoU/entities"
	"github.com/joaopedrosgs/OpenLoU/hub"
	"github.com/joaopedrosgs/OpenLoU/mapserver"
	"github.com/joaopedrosgs/OpenLoU/session"

	"github.com/joaopedrosgs/OpenLoU/test"
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/joaopedrosgs/OpenLoU/database"
)

var context = log.WithField("Entity", "OpenLoU")

func main() {

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	context.Info("OpenLoU is starting...")
	database.InitDB()
	entities.RegisterAllTroops()
	entities.RegisterAllConstructions()

	Sessions := session.NewSessionInMemory()
	Hermes := hub.Create(Sessions)

	LoginServer, err := accountserver.New(Sessions)
	if err != nil {
		context.Error(err.Error())
	}

	MapServer := mapserver.New()
	CityServer := cityserver.New()

	err = Hermes.RegisterWorker(CityServer)
	if err != nil {
		context.Error(err.Error())
	}
	err = Hermes.RegisterWorker(MapServer)
	if err != nil {
		context.Error(err.Error())
	}

	go MapServer.StartListening()
	go LoginServer.StartListening(":8000")
	go func() {
		time.Sleep(4 * time.Second)
		for i := 0; i < 10; i++ {
			go test.RunTest()
		}

	}()
	Hermes.StartListening()

}
