package database

import (
	"errors"

	"github.com/jackc/pgx"

	"github.com/joaopedrosgs/OpenLoU/configuration"

	log "github.com/sirupsen/logrus"
)

var db *pgx.ConnPool
var context = log.WithField("Entity", "Database")

func InitDB() {
	err := errors.New("")

	db, err = pgx.NewConnPool(pgx.ConnPoolConfig{ConnConfig: configuration.GetConnConfig(), MaxConnections: 5})
	if err != nil {
		context.WithField("When", "Database init").Error(err.Error())
	}
	var numberOfContinent uint = 0
	expectedNumberOfContinents := configuration.GetSingleton().Parameters.General.WorldSize
	db.QueryRow("Select Count(*) from continents").Scan(&numberOfContinent)
	if numberOfContinent < expectedNumberOfContinents {
		context.Warning("Looks like there are less continents than expected, attempting to create more...")
		createNewContinents()
	}

}

func Close() {
	db.Close()
}
