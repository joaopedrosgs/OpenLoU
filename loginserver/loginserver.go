package loginserver

import (
	"OpenLoU/configuration"
	"database/sql"
	"errors"
	"fmt"

	"OpenLoU/communication"
	_ "github.com/lib/pq" // Postgresql Driver
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// Const values
const (
	KeySize    = 64
	DBSessions = "sessions"
	DBUsers    = "users"

	//Queries
	newUserQuery           = "INSERT INTO " + DBUsers + "(login, password, email) VALUES ($1, $2, $3)  RETURNING id"
	loginQuery             = "SELECT id, login, password FROM " + DBUsers + " WHERE login =  $1 LIMIT 1 "
	userExistsQuery        = "SELECT 1 FROM " + DBUsers + " WHERE login=$1 LIMIT 1"
	deleteUserByLoginQuery = "DELETE from " + DBUsers + " WHERE login=$1"
	newSessionQuery        = "INSERT INTO " + DBSessions + " VALUES($1, $2, $3)"
	findSessionQuery       = "SELECT 1 FROM " + DBSessions + " WHERE key=$1 AND user_id=$2 AND ip=$3 LIMIT 1"
	connectionString       = "user=%s password=%s host=%s port=%d dbname=%s sslmode=%s"
)

type LoginServer struct {
	Database *sql.DB
	In       chan communication.Request
	Out      chan communication.Answer
}

type LoginAttempt struct {
	IP       string
	Login    string
	Password string
}

type Answer struct {
	Auth bool
	Key  string
}

func (s *LoginServer) StartListening() {
}

// CreateAndConnect returns an LoginServer that deals with the authentication of the user
func CreateAndConnect(config *configuration.Config) (*LoginServer, error) {

	connectionString := fmt.Sprintf(connectionString, config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Name, config.Db.SSL)
	database, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.WithFields(log.Fields{"Error": err.Error()}).Fatal(dbError)
		return nil, err

	}
	err = database.Ping()
	if err != nil {
		log.WithFields(log.Fields{"Error": err.Error()}).Fatal(dbError)
		return nil, err
	}
	log.WithFields(log.Fields{"From": "Login Server"}).Info("Database connection established")

	return &LoginServer{database, make(chan communication.Request), make(chan communication.Answer)}, nil

}

//NewAttempt returns an Answer which contains the auth info from the attempt
func (s *LoginServer) NewAttempt(info LoginAttempt) (answer Answer) {

	user, err := s.CheckCredentials(info)
	if err != nil {
		answer.Auth = false
		log.WithFields(log.Fields{userWrongIP: info.IP, "Error": err.Error()}).Info("Login Server")
	} else {
		key := GenUniqueKey(KeySize)
		s.CreateSession(user, key, info.IP)
		answer = Answer{true, key}
		log.WithFields(log.Fields{userLogged: info.Login}).Info("Login Server")
	}
	return
}

//CheckCredentials returns the user and nil if the credentials are correct
func (s *LoginServer) CheckCredentials(attempt LoginAttempt) (*user, error) {
	if len(attempt.Password) == 0 || len(attempt.Login) == 0 || len(attempt.IP) == 0 {
		return nil, errors.New(emptyFields)
	}
	user := user{}
	var pass string
	err := s.Database.QueryRow(loginQuery, attempt.Login).Scan(&user.Id, &user.Login, &pass)
	if err != nil {
		return nil, errors.New(accountInexistent)
	}
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(attempt.Password))
	if err != nil {
		return nil, errors.New(wrongPass + err.Error())
	}
	return &user, nil
}
