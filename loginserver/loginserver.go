package loginserver

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Const values
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

// Const errors
const (
	EmptyFields       = "Empty fields"
	ShortPassword     = "Password is too short"
	DBError           = "Failed to access database"
	AccountInexistent = "Account not found"
	AccountExists     = "An account already exists with that information"
)

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
		log.WithFields(log.Fields{"user failed ip": info.Ip, "Error": err.Error()}).Info("Login Server")
	} else {
		key := GenUniqueKey(KeySize)
		s.CreateSession(user, key, info.Ip)
		answer = Answer{true, key}
		log.WithFields(log.Fields{"user logged in": info.Login}).Info("Login Server")
	}
	return
}

//CheckCredentials returns nil if the credentials are correct
func (s *loginServer) CheckCredentials(attempt LoginAttempt) (error, *user) {
	if len(attempt.Password) == 0 || len(attempt.Login) == 0 || len(attempt.Ip) == 0 {
		return errors.New("Empty fields"), nil
	} else {
		db := s.Database.Copy().DB(DBName).C(DBUsers)
		user := user{}
		err := db.Find(bson.M{"login": attempt.Login}).One(&user)
		if err != nil {
			return errors.New("Account doesn't exist"), nil
		}
		err = bcrypt.CompareHashAndPassword(user.Pass, []byte(attempt.Password))
		if err != nil {
			return errors.New("Wrong password: " + err.Error()), nil
		}
		return nil, &user

	}
}
