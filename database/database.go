package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/entities"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB
var context = log.WithField("Entity", "Database")

func InitDB() {
	err := errors.New("")
	db, err = gorm.Open("postgres", configuration.GetConnectionString())
	if err != nil {
		context.Error(err.Error())
	}
	db.AutoMigrate(&entities.User{}, &entities.City{}, &entities.Construction{}, &entities.Upgrade{})

}

func GetSingleton() *gorm.DB {
	if db == nil {
		InitDB()
	}
	return db
}
