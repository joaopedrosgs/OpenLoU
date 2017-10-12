package loginserver

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joaopedrosgs/OpenLoU/configuration"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// Const values
const (
	KeySize    = 64
	DBSessions = "sessions"
	DBUsers    = "users"

	//Queries
	newUserQuery      = "INSERT INTO " + DBUsers + "(login, password, email) VALUES ($1, $2, $3)  RETURNING id"
	loginQuery        = "SELECT id, login, password FROM " + DBUsers + " WHERE login =  $1 LIMIT 1 "
	userExists        = "SELECT 1 FROM " + DBUsers + " WHERE login=$1 LIMIT 1"
	deleteUserByLogin = "DELETE from " + DBUsers + " WHERE login=$1"
	newSessionQuery   = "INSERT INTO " + DBSessions + " VALUES($1, $2, $3)"
	findSessionQuery  = "SELECT 1 FROM " + DBSessions + " WHERE key=$1 AND user_id=$2 AND ip=$3 LIMIT 1"
)

type server struct {
	Database *sql.DB
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

func (s *server) StartListening() {
}

// New returns an server that deals with the authentication of the user
func New(isDebug bool, config *configuration.Config) *server {
	log.SetOutput(os.Stdout)

	if isDebug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d "+
		"dbname=%s sslmode=%s", config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Name, config.Db.SSL)
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.WithFields(log.Fields{"Error": err.Error()}).Fatal(dbError)
		return nil

	}
	err = database.Ping()
	if err != nil {
		log.WithFields(log.Fields{"Error": err.Error()}).Fatal(dbError)
		return nil
	}
	log.WithFields(log.Fields{"From": "Login Server"}).Info("Database connection established")

	return &server{database}

}

//NewAttempt returns an Answer which contains the auth info from the attempt
func (s *server) NewAttempt(info LoginAttempt) (answer Answer) {

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
func (s *server) CheckCredentials(attempt LoginAttempt) (*user, error) {
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
