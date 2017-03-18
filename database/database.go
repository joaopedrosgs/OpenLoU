package database

import (
	"fmt"

	"LordOfUltima/configuration"

	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

var Session *mgo.Session

func Open() {
	var err error
	println("Conectando ao banco de dados")
	Session, err = mgo.Dial(configuration.Parameters.Connection)
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)

	if err != nil {
		fmt.Printf("Falha ao conectar no banco:", err)
	}
	println("Conectado ao banco de dados!")
}
