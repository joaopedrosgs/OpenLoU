package storage

import (
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
	"github.com/joaopedrosgs/OpenLoU/session"
)

func CreateUser(db *pgx.Conn, login, passwordHash, email string) error {

	var userID int

	err := db.QueryRow("INSERT INTO users (name, email, password) values ($1, $2, $3) returning id", login, email, passwordHash).Scan(&userID)

	if err != nil {
		logger.Error("account could not be created because: " + err.Error())
		return err
	}
	return nil
}

func GetUserInfo(db *pgx.Conn, userName string) (*models.User, error) {
	user := &models.User{}
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
		logger.WithField("When", "Retrieving user information").Error(err.Error())
		return nil, err
	}
	return user, nil
}

func GetUserInfoByKey(db *pgx.Conn, key string) (*models.User, error) {
	userName, err := session.GetUserName(key)
	if err != nil {
		logger.Error("failed to get user name by key: ", err.Error())
	}
	user, err := GetUserInfo(db, userName)
	if err != nil {
		logger.Error("failed to get user info: ", err.Error())
	}
	return user, nil

}
