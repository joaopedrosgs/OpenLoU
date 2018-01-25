package database

import (
	"errors"
	"github.com/joaopedrosgs/OpenLoU/entities"
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
	randX, randY, continentID, err := findNewCityLocation()
	if err != nil {
		context.WithField("When", "Finding location").Error(err.Error())
		return err
	}
	err = createCity(userID, randX, randY, continentID)
	if err != nil {
		context.WithField("When", "Creating city").Error(err.Error())
		return err
	}
	return nil
}

func GetUser(email string) (*entities.User, error) {
	if db == nil {
		InitDB()
	}
	user := &entities.User{}
	err := db.QueryRow("SELECT * from users where email = $1", email).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, errors.New("Account could not be found: " + err.Error())
	}
	return user, nil
}
