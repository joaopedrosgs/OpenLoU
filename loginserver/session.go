package loginserver

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
)

type Session struct {
	UID        int
	Ip         string
	Key        string
	LoggedIn   time.Time
	LastAction time.Time
}

//CreateSessions returns nil if it was able to connect to the database and store the session created
func (s *LoginServer) CreateSession(user *user, key string, ip string) error {
	if user == nil || len(key) == 0 || len(ip) == 0 {
		return errors.New(emptyFields)
	}
	newsession := Session{LoggedIn: time.Now(), LastAction: time.Now(), UID: user.Id, Key: key, Ip: ip}
	_, err := s.Database.Exec(newSessionQuery, newsession.Key, newsession.UID, newsession.Ip)
	if err != nil {
		log.WithFields(log.Fields{"Failed to create session to": user.Login}).Info("Login Server")
	}
	return err
}

//SessionExists returns an error if session doesn't exist
func (s *LoginServer) SessionExists(session Session) error {
	res := 0
	err := s.Database.QueryRow(findSessionQuery, session.Key, session.UID, session.Ip).Scan(&res)
	if err != nil {
		return err
	}
	if res == 0 {
		return errors.New("Session could not be found")
	}
	return nil

}
