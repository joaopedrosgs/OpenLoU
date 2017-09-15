package main

import (
	"./gameEntities"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"time"
)

var loadedConstructions gameEntities.ConstructionsMap

type Session struct {
	LoggedIn   time.Time
	LastAction time.Time
	Login      string
	Key        string
}

type LoginServer struct {
	/*Db         *database*/
	MinAllowed byte
	Sessions   map[string]Session
}
type LoginAttempt struct {
	Ip       string
	Login    string
	Password string
}

const (
	KeySize = 64
)

func (s *LoginServer) StartAndListen(port int) {
	s.Sessions = make(map[string]Session)
}

func (s *LoginServer) NewAttempt(info LoginAttempt) (answer Answer) {
	err := CheckCredentials(info)
	if err != nil {
		log.WithField("Login-Server", err)
	} else {
		key := genUniqueKey(KeySize)
		s.CreateSession(info, key)
		body, _ := json.Marshal(key)

		answer = Answer{Ok, LoggedIn, string(body)}
	}
	return
}
func (s *LoginServer) CreateSession(attempt LoginAttempt, key string) {
	s.Sessions[attempt.Login] = Session{time.Now(), time.Now(), attempt.Login, key}

}
func CheckCredentials(attempt LoginAttempt) error {
	if len(attempt.Password) == 0 || len(attempt.Login) == 0 || len(attempt.Ip) == 0 {
		return errors.New("Empty fields")
	} else {
		if attempt.Login == "login" && attempt.Password == "pass" { //to do
			return nil
		} else {
			return errors.New("Account not found")
		}
	}
}
