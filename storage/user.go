package storage

import (
	"errors"
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *pgx.Conn, login, password, email string) error {

	if len(login) < 6 || len(email) < 8 || len(password) < 8 {
		return errors.New("login, email or password is too short")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return errors.New("internal error")
	}
	_, err = db.Exec("INSERT INTO users (name, email, password) values ($1, $2, $3)", login, email, passwordHash)

	if err != nil {
		logger.Error("account could not be created because: " + err.Error())
		return err
	}
	return nil
}

func GetUserInfo(db *pgx.Conn, userName string) (*models.User, error) {
	user := &models.User{}
	err := db.QueryRow(
		"SELECT name, email, alliance_name, gold, diamonds, darkwood, runestone, veritium, trueseed, rank, password "+
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
		&user.Rank,
		&user.PasswordHash)
	if err != nil {
		logger.WithField("When", "Retrieving user information").Error(err.Error())
		return nil, err
	}
	return user, nil
}
