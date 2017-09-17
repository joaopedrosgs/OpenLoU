package loginserver

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
)

type Auth struct {
	Login string
	Key   string
	Ip    string
}
type Session struct {
	LoggedIn   time.Time
	LastAction time.Time
	AuthInfo   Auth
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

type Answer struct {
	Auth bool
	Key  string
}

func (s *LoginServer) NewAttempt(info LoginAttempt) (answer Answer) {

	err := CheckCredentials(info)
	if err != nil {
		answer.Auth = false
		log.WithFields(log.Fields{"User failed ip": info.Ip, "Error": err.Error()}).Info("Login Server")
	} else {
		key := genUniqueKey(KeySize)
		s.CreateSession(info, key)
		answer = Answer{true, key}
		log.WithFields(log.Fields{"User logged in": info.Login}).Info("Login Server")
	}
	return
}
func (s *LoginServer) CreateSession(attempt LoginAttempt, key string) {
	s.Sessions[key] = Session{time.Now(), time.Now(), Auth{attempt.Login, key, attempt.Ip}}

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

func (s *LoginServer) SessionExists(authinfo Auth) error {
	if sessionStored, ok := s.Sessions[authinfo.Key]; ok {
		if sessionStored.AuthInfo.Ip == authinfo.Ip {
			return nil
		}
		return errors.New("Wrong IP")
	}
	return errors.New("Session could not be found")

}
