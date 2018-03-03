package database

import (
	"errors"
	"github.com/joaopedrosgs/OpenLoU/entities"
	"github.com/joaopedrosgs/OpenLoU/session"
)

func CreateUser(login, passwordHash, email string) error {
	if db == nil {
		InitDB()
	}
	var userID uint

	err := db.QueryRow("INSERT INTO users (name, email, password) values ($1, $2, $3) returning id", login, email, passwordHash).Scan(&userID)

	if err != nil {
		return errors.New("Account could not be created because: " + err.Error())
	}
	available, err := checkContinentSpace()
	if err != nil {
		context.WithField("When", "Creating city").Error(err.Error())
		return err
	}
	if !available {
		err := startingNewContinent()
		if err != nil {
			context.WithField("When", "Finding location").Error(err.Error())
			return err
		}
	}
	randX, randY, continentX, continentY, err := findNewCityLocation()
	if err != nil {
		context.WithField("When", "Finding location").Error(err.Error())
		return err
	}
	err = createCity(entities.City{
		TileNode: entities.TileNode{randX, randY, continentX, continentY, "city"},
		Name:     "New City",
		Points:   3,
		UserName: login,
	})
	if err != nil {
		context.WithField("When", "Creating city").Error(err.Error())
		return err
	}
	return nil
}

func GetUserInfo(userName string) (*entities.User, error) {
	user := &entities.User{}
	err := db.QueryRow(
		"SELECT name, email, alliance_id, gold, diamonds, darkwood, runestone, veritium, trueseed, rank "+
			"from users "+
			"WHERE name = $1", userName).Scan(
		&user.Name,
		&user.Email,
		&user.AllianceName,
		&user.Gold,
		&user.Diamonds,
		&user.Darkwood,
		&user.Runestone,
		&user.Veritium,
		&user.Trueseed,
		&user.Rank)
	if err != nil {
		context.WithField("When", "Retrieving user information").Error(err.Error())

	}
	return user, err
}

func GetUserInfoByKey(key string) (*entities.User, error) {
	userName, ok := session.GetUserName(key)
	if ok {
		return GetUserInfo(userName)
	} else {
		return nil, errors.New("Account not found")
	}
}
