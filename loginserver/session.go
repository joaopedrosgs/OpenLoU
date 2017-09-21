package loginserver

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct {
	Id         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	LoggedIn   time.Time
	LastAction time.Time
	Login      string
	Ip         string
	Key        string
}

//CreateSessions returns nil if it was able to connect to the database and store the session created
func (s *loginServer) CreateSession(user *user, key string, ip string) error {
	newsession := Session{LoggedIn: time.Now(), LastAction: time.Now(), Login: user.Login, Key: key, Ip: ip}
	db := s.Database.Copy()
	defer db.Close()
	sessions := db.DB(DBName).C(DBSessions)
	n, _ := sessions.Find(bson.M{"login": user.Login}).Count()
	if n >= SessionLimit {
		i := sessions.Find(bson.M{"login": user.Login}).Sort("lastaction").Limit(n - (SessionLimit - 1)).Iter()
		ele := Session{}
		for i.Next(&ele) {
			sessions.RemoveId(ele.Id)
		}
	}
	err := sessions.Insert(newsession)
	if err != nil {
		log.WithFields(log.Fields{"Failed to create session to": user.Login}).Info("Login Server")
	}
	return err
}

//SessionExists returns an error if session doesn't exist
func (s *loginServer) SessionExists(session Session) error {
	sessionStored := Session{}
	db := s.Database.Copy()
	defer db.Close()
	err := db.DB(DBName).C(DBSessions).Find(bson.M{"key": session.Key}).One(&sessionStored)

	if err == nil {
		if sessionStored.Ip == session.Ip {
			return nil
		}
		return errors.New("Wrong IP")
	}
	return errors.New("Session could not be found")

}
