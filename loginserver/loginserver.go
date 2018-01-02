package loginserver

import (
	"OpenLoU/configuration"
	"database/sql"
	"errors"
	"fmt"

	"OpenLoU/communication"
	"OpenLoU/hermes"
	"encoding/json"
	_ "github.com/lib/pq" // Postgresql Driver
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"net/http"
)

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

func (s *LoginServer) StartListening() {
	// Index Handler
	http.HandleFunc("/login", s.loginHandler)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
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
		log.WithFields(log.Fields{"Error": err.Error()}).Fatal(dbError)
		return nil, err

	}
	log.WithFields(log.Fields{"From": "Login Server"}).Info("Database connection established")

	return &LoginServer{database, backend}, nil

}

//NewAttempt returns an Answer which contains the auth info from the attempt
func (s *LoginServer) NewAttempt(attempt *LoginAttempt) *communication.Answer {
	answer := &communication.Answer{}
	id, err := s.CheckCredentials(attempt)

	if err != nil {
		answer.Result = false
		answer.Data = err.Error()
	} else {

		key := GenUniqueKey()
		answer.Result = true
		answer.Data = key
		answer.Type = NEW_LOGIN
		s.sessions.NewSession(id, key)
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
	if err != nil {
		return -1, errors.New(accountInexistent)
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(attempt.Password))
	if err != nil {
		return -1, errors.New(wrongPass)
	}
	return id, nil
}
