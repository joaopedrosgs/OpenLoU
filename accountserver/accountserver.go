package accountserver

import (
	"errors"
	"fmt"

	"github.com/joaopedrosgs/OpenLoU/configuration"

	"encoding/json"
	"net/http"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/entities"
	"github.com/joaopedrosgs/OpenLoU/session"
	_ "github.com/lib/pq" // Postgresql Driver
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"time"
)

// Const values
const (

	//ANSWER RETURN VALUES

	NEW_LOGIN = 301

	///////////////////////////
	DBUsers = "users"

	//Queries
	newUserQuery           = "INSERT INTO " + DBUsers + "(name, password_hash, email) VALUES ($1, $2, $3)  RETURNING id"
	loginQuery             = "SELECT password_hash, id FROM " + DBUsers + " WHERE email =  $1 LIMIT 1 "
	userExistsQuery        = "SELECT 1 FROM " + DBUsers + " WHERE email=$1 LIMIT 1"
	deleteUserByLoginQuery = "DELETE from " + DBUsers + " WHERE email=$1"
)

type LoginServer struct {
	sessions *session.SessionMem
	context  *log.Entry
}

type LoginAttempt struct {
	Email    string
	Password string
}

func New(backend *session.SessionMem) (*LoginServer, error) {
	return &LoginServer{backend, log.WithFields(log.Fields{"Entity": "Account Server"})}, nil

}

func (s *LoginServer) StartListening(address string) {
	// Index Handler
	http.HandleFunc("/login", s.loginHandler)
	http.HandleFunc("/register", s.registerHandler)
	err := http.ListenAndServe(address, nil)
	for err != nil {
		s.context.Error("Failed to listen: " + err.Error())
		s.context.Info("Trying again in 10 seconds...")
		time.Sleep(10 * time.Second)
		err = http.ListenAndServe(address, nil)

	}
	s.context.Info("Account server has started listening")
}
func (s *LoginServer) loginHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {
		email := request.PostFormValue("email")
		password := request.PostFormValue("password")
		attempt := &LoginAttempt{email, password}
		answer := s.NewAttempt(attempt)
		jsonAnswer, _ := json.Marshal(answer)
		fmt.Fprintf(writer, string(jsonAnswer))

	}

}
func (s *LoginServer) registerHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		login := request.PostFormValue("login")
		email := request.PostFormValue("email")
		password := request.PostFormValue("password")
		answer := s.CreateAccount(login, email, password)
		jsonAnswer, _ := json.Marshal(answer)
		fmt.Fprintf(writer, string(jsonAnswer))
	}
}

// New returns an LoginServer that deals with the authentication of the user

//NewAttempt returns an Answer which contains the auth info from the attempt
func (s *LoginServer) NewAttempt(attempt *LoginAttempt) *communication.Answer {
	answer := &communication.Answer{}
	id, err := s.CheckCredentials(attempt)

	if err != nil {
		answer.Ok = false
		answer.Data = err.Error()
	} else {

		key, err := GenerateRandomString(configuration.GetSingleton().Parameters.Security.KeySize)
		if err == nil {
			created := s.sessions.NewSession(id, key)
			if created {
				answer.Ok = true
				answer.Data = key
				answer.Type = NEW_LOGIN
			} else {
				answer.Data = "Failed to create session"
			}
		}
	}

	return answer
}

//CheckCredentials returns the user and nil if the credentials are correct
func (s *LoginServer) CheckCredentials(attempt *LoginAttempt) (uint, error) {
	if len(attempt.Password) == 0 || len(attempt.Email) == 0 {
		return 0, errors.New(emptyFields)
	}

	user := entities.User{}
	database.GetSingleton().Where("email = ?", attempt.Email).First(&user)
	database.GetSingleton().Preload("Cities").Find(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(attempt.Password))
	if err != nil {
		return 0, errors.New(wrongPass)
	}
	return user.ID, nil
}
func (s *LoginServer) CreateAccount(login string, email string, password string) *communication.Answer {
	answer := &communication.Answer{}

	if len(login) < 6 || len(email) < 8 || len(password) < 8 {
		answer.Data = "Too small"
		return answer
	}
	password_hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {

		user := entities.User{Name: login, Email: email, PasswordHash: string(password_hash)}
		database.GetSingleton().NewRecord(user)
		database.GetSingleton().Create(&user)

		//go s.initUserAccount(&user)
		answer.Data = "Success"
		answer.Ok = true
	}

	return answer

}
func (server *LoginServer) initUserAccount(i int) {

}
