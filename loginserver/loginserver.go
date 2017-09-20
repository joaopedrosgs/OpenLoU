package loginserver

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	KeySize      = 64
	DBName       = "OpenLoU"
	DBSessions   = "sessions"
	DBUsers      = "users"
	DBPort       = 27017
	DBUsername   = ""
	DBPassword   = ""
	DBTimeout    = 60 * time.Second
	DBSsh        = true
	SessionLimit = 3
)

var DBAddress = []string{"127.0.0.1"}

type loginServer struct {
	Database *mgo.Session
}

func (server *loginServer) StartListening() {

}

type LoginAttempt struct {
	Ip       string
	Login    string
	Password string
}

type Answer struct {
	Auth bool
	Key  string
}

// New returns an loginServer that deals with the authentication of the user
func New(maxSessions int) *loginServer {
	dialInfo := &mgo.DialInfo{
		Addrs:    DBAddress,
		Timeout:  DBTimeout,
		Database: DBName,
		Username: DBUsername,
		Password: DBPassword,
	}
	database, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.WithFields(log.Fields{"Error": err.Error()}).Fatal("Couldn't connect to the login database")

	} else {
		database.SetMode(mgo.Monotonic, true)
	}
	return &loginServer{database}
}

//NewAttempt returns an Answer which contains the auth info from the attempt
func (s *loginServer) NewAttempt(info LoginAttempt) (answer Answer) {

	err, user := s.CheckCredentials(info)
	if err != nil {
		answer.Auth = false
		log.WithFields(log.Fields{"User failed ip": info.Ip, "Error": err.Error()}).Info("Login Server")
	} else {
		key := genUniqueKey(KeySize)
		s.CreateSession(user, key, info.Ip)
		answer = Answer{true, key}
		log.WithFields(log.Fields{"User logged in": info.Login}).Info("Login Server")
	}
	return
}

//CreateSessions returns nil if it was able to connect to the database and store the session created
func (s *loginServer) CreateSession(user *User, key string, ip string) error {
	newsession := Session{LoggedIn: time.Now(), LastAction: time.Now(), Login: user.Login, Key: key, Ip: ip}
	db := s.Database.Copy()
	defer db.Close()
	sessions := db.DB(DBName).C(DBSessions)
	n, _ := sessions.Find(bson.M{"login": user.Login}).Count()
	if n >= SessionLimit {
		i := sessions.Find(bson.M{"login": user.Login}).Sort("lastaction").Limit(n - (SessionLimit - 1)).Iter()
		ele := User{}
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

//CheckCredentials returns nil if the credentials are correct
func (s *loginServer) CheckCredentials(attempt LoginAttempt) (error, *User) {
	if len(attempt.Password) == 0 || len(attempt.Login) == 0 || len(attempt.Ip) == 0 {
		return errors.New("Empty fields"), nil
	} else {
		db := s.Database.Copy().DB(DBName).C(DBUsers)
		user := User{}
		err := db.Find(bson.M{"login": attempt.Login}).One(&user)
		if err != nil {
			return errors.New("Account doesn't exist"), nil
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(attempt.Password))
		if err != nil {
			return errors.New("Wrong password"), nil
		}
		return nil, &user

	}
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
