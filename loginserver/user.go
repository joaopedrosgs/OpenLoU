package loginserver

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

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

func (s *LoginServer) NewUser(login, pass, email string) (*user, error) {

	if len(login) == 0 || len(email) == 0 {
		return nil, errors.New(emptyFields)
	}
	if len(pass) < MinPassSize {
		return nil, errors.New(shortPassword)

	}
	hashpass, err := bcrypt.GenerateFromPassword([]byte(pass), 4)
	if err != nil {
		return nil, errors.New("There was a problem with the hash function")
	}
	user := user{Login: login, Email: email, Created: time.Now()}
	id := -1
	err = s.Database.QueryRow(newUserQuery, user.Login, hashpass, user.Email).Scan(&id)
	if err != nil {
		return nil, err
	}
	user.Id = id
	return &user, err
}

func (s *LoginServer) LoadUserByLogin(login string) (*user, error) {
	if len(login) == 0 {
		return nil, errors.New(emptyFields)
	}
	u := user{}
	err := s.Database.QueryRow(loginQuery, login).Scan(&u.Id, &u.Login)
	return &u, err
}

func (l *LoginServer) SaveUserChanges(u *user) error {
	return nil
}

func (s *LoginServer) UserExists(login, email string) error {
	res := 0
	err := s.Database.QueryRow(userExistsQuery, login).Scan(&res)
	if err != nil {
		return errors.New(dbError)
	}
	if res == 0 {
		return errors.New(accountInexistent)
	}
	return nil

}
func (s *LoginServer) DeleteUserByLogin(login string) error {
	_, err := s.Database.Exec(deleteUserByLoginQuery, login)
	return err
}
