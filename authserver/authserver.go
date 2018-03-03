package authserver

import (
	"errors"
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/session"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"github.com/joaopedrosgs/OpenLoU/entities"

	"time"
)

type authServer struct {
	context *log.Entry
}

type LoginAttempt struct {
	Login    string
	Password string
}

func New() (*authServer, error) {
	return &authServer{log.WithFields(log.Fields{"Entity": "Auth Server"})}, nil

}

func (s *authServer) StartListening(address string) {
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
	s.context.Info("Auth server has started listening")
}
func (s *authServer) loginHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		fmt.Fprintf(writer, notPost)

	} else {
		answer := communication.Answer{}
		email := request.PostFormValue("email")
		password := request.PostFormValue("password")
		attempt := &LoginAttempt{email, password}
		if key, err := s.NewAttempt(attempt); err != nil {
			answer.Data = err.Error()
		} else {
			answer.Ok = true
			answer.Data = key
		}
		jsonAnswer, err := json.Marshal(answer)
		if err != nil {
			jsonAnswer = []byte(InternalError)
			s.context.WithField("When", "Converting login answer to JSON").Error(err.Error())
		}
		fmt.Fprintf(writer, string(jsonAnswer))

	}

}
func (s *authServer) registerHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		fmt.Fprint(writer, notPost)
	} else {
		login := request.PostFormValue("login")
		email := request.PostFormValue("email")
		password := request.PostFormValue("password")
		answer := s.createAccount(login, email, password)
		jsonAnswer, _ := json.Marshal(answer)
		fmt.Fprint(writer, string(jsonAnswer))
	}
} /*
func (s *authServer) userInfoHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		fmt.Fprint(writer, notPost)
	} else if session, exists := session.Exists(request.PostFormValue("key")); !exists {
		fmt.Fprint(writer, sessionNotFound)
	} else if userInfo, err := database.GetUserInfo(session.UserName); err != nil {
		s.context.WithField("When", "Retrieving user information").Error(err.Error())
		fmt.Fprint(writer, InternalError)
	} else if jsonUserInfo, err := json.Marshal(userInfo); err != nil {
		s.context.WithField("When", "Converting answer to json").Error(err.Error())
		fmt.Fprintf(writer, InternalError)
	} else {
		fmt.Fprint(writer, jsonUserInfo)
	}

}*/

//NewAttempt returns an Answer which contains the auth info from the attempt
func (s *authServer) NewAttempt(attempt *LoginAttempt) (string, error) {
	user, err := s.checkCredentials(attempt)
	payload := ""
	if err != nil {
		payload = err.Error()
	} else if key, err := session.NewSession(user); err != nil {
		payload = err.Error()
	} else {
		payload = key
	}

	return payload, err
}

//checkCredentials returns the user and nil if the credentials are correct
func (s *authServer) checkCredentials(attempt *LoginAttempt) (entities.User, error) {
	var err error
	var user entities.User
	if len(attempt.Password) < 8 || len(attempt.Login) < 8 {
		err = errors.New(emptyFields)
	} else if user, err := database.GetUserInfo(attempt.Login); err != nil {
		err = errors.New(wrongAccountInfo)
	} else if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(attempt.Password)); err != nil {
		err = errors.New(wrongAccountInfo)
	}
	return user, err

}
func (s *authServer) createAccount(login string, email string, password string) *communication.Answer {
	answer := &communication.Answer{}
	if len(login) < 6 || len(email) < 8 || len(password) < 8 {
		answer.Data = shortCredentials
	} else if passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10); err != nil {
		answer.Data = InternalError
	} else if err = database.CreateUser(login, string(passwordHash), email); err != nil {
		answer.Data = accountExists + err.Error()
	} else {
		answer.Data = success
		answer.Ok = true
	}
	return answer

}
