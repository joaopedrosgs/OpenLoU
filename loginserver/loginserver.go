package loginserver

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/joaopedrosgs/OpenLoU/configuration"

	"encoding/json"
	"net/http"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/hermes"
	_ "github.com/lib/pq" // Postgresql Driver
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var context = log.WithFields(log.Fields{"Entity": "Login Server"})

// Const values
const (

	//ANSWER RETURN VALUES

	NEW_LOGIN = 301

	///////////////////////////
	DBUsers = "users"

	//Queries
	newUserQuery           = "INSERT INTO " + DBUsers + "(login, password, email) VALUES ($1, $2, $3)  RETURNING id"
	loginQuery             = "SELECT password, id FROM " + DBUsers + " WHERE email =  $1 LIMIT 1 "
	userExistsQuery        = "SELECT 1 FROM " + DBUsers + " WHERE email=$1 LIMIT 1"
	deleteUserByLoginQuery = "DELETE from " + DBUsers + " WHERE email=$1"
)

type LoginServer struct {
	Database *sql.DB
	sessions hermes.ISessionBackend
}

type LoginAttempt struct {
	Email    string
	Password string
}

type Answer struct {
	Auth bool
	Key  string
}

func (s *LoginServer) StartListening(address string) {
	// Index Handler
	http.HandleFunc("/login", s.loginHandler)
	err := http.ListenAndServe(address, nil)
	for err != nil {
		context.Error("Failed to listen: " + err.Error())
		context.Info("Trying again in 10 seconds...")
		time.Sleep(10 * time.Second)
		err = http.ListenAndServe(address, nil)

	}
	context.Info("Login Server has started listening")
}
func (s *LoginServer) loginHandler(writer http.ResponseWriter, request *http.Request) {
	answer := communication.BadRequest()

	if request.Method == "POST" {
		jsonRequest := request.PostFormValue("data")
		attempt := &LoginAttempt{}
		err := json.Unmarshal([]byte(jsonRequest), &attempt)
		if err == nil {
			answer = s.NewAttempt(attempt)
		}
	}

	json, _ := json.Marshal(answer)
	fmt.Fprintf(writer, string(json))

}

// New returns an LoginServer that deals with the authentication of the user
func New(backend hermes.ISessionBackend) (*LoginServer, error) {
	database, err := sql.Open("postgres", configuration.GetConnectionString())

	if err != nil {
		context.Error(err.Error())
		return nil, err

	}
	context.Info("Login server database connection established")

	return &LoginServer{database, backend}, nil

}

//NewAttempt returns an Answer which contains the auth info from the attempt
func (s *LoginServer) NewAttempt(attempt *LoginAttempt) *communication.Answer {
	answer := &communication.Answer{}
	id, err := s.CheckCredentials(attempt)

	if err != nil {
		answer.Ok = false
		answer.Data = err.Error()
	} else {

		key, err := GenerateRandomString(configuration.GetInstance().Parameters.Security.KeySize)
		if err == nil {

			answer.Ok = true
			answer.Data = key
			answer.Type = NEW_LOGIN
			s.sessions.NewSession(id, key)
		}
	}

	return answer
}

//CheckCredentials returns the user and nil if the credentials are correct
func (s *LoginServer) CheckCredentials(attempt *LoginAttempt) (int, error) {
	if len(attempt.Password) == 0 || len(attempt.Email) == 0 {
		return -1, errors.New(emptyFields)
	}

	pass := ""
	id := -1
	err := s.Database.QueryRow(loginQuery, attempt.Email).Scan(&pass, &id)
	if id < 0 {
		return -1, errors.New(accountInexistent)
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(attempt.Password))
	if err != nil {
		return -1, errors.New(wrongPass)
	}
	return id, nil
}
