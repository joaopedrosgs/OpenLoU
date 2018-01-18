package accountserver

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/entities"
	"net"
	"time"
)

const MinPassSize = 7

type UserConnInfo struct {
	Connection *net.Conn
	UserID     int
}

type user struct {
	Id         int
	Login      string
	Email      string
	Created    time.Time
	LastLogin  time.Time
	CitiesId   []int
	SessionsId []int
}

func (s *LoginServer) NewUser(login, pass, email string) error {

	if len(login) == 0 || len(email) == 0 {
		return errors.New(emptyFields)
	}
	if len(pass) < MinPassSize {
		return errors.New(shortPassword)

	}
	hashpass, err := bcrypt.GenerateFromPassword([]byte(pass), 4)
	if err != nil {
		return errors.New("There was a problem with the hash function")
	}
	user := entities.User{Name: login, PasswordHash: string(hashpass), Email: email}
	ok := database.GetSingleton().NewRecord(user) // => returns `true` as primary key is blank
	if !ok {
		return errors.New("Error creating account")
	}
	database.GetSingleton().Create(&user)
	return err
}

func (l *LoginServer) SaveUserChanges(u *user) error {
	return nil
}
