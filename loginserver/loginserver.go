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
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
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

	if err := http.ListenAndServe(":12345", nil); err != nil {
		println(err.Error())
	}
	println("Login server has been started")
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
		println(err.Error())
		return nil, err

	}
	println("Login server: Database connection established")

	return &LoginServer{database, backend}, nil

}

//NewAttempt returns an Answer which contains the auth info from the attempt
func (s *LoginServer) NewAttempt(attempt *LoginAttempt) *communication.Answer {
	answer := &communication.Answer{}
	answer.Data = make(map[string]string)
	id, err := s.CheckCredentials(attempt)

	if err != nil {
		answer.Data["Result"] = "False"
		answer.Data["Message"] = err.Error()
	} else {

		key := GenUniqueKey()
		answer.Data["Result"] = "True"
		answer.Data["Key"] = key
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
	if id < 0 {
		return -1, errors.New(accountInexistent)
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(attempt.Password))
	if err != nil {
		return -1, errors.New(wrongPass)
	}
	return id, nil
}
