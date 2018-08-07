package authserver

import (
	"errors"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/server"
	"github.com/joaopedrosgs/OpenLoU/session"
	"github.com/joaopedrosgs/OpenLoU/storage"
	"golang.org/x/crypto/bcrypt"

	"github.com/joaopedrosgs/OpenLoU/models"
)

type authServer struct {
	server.Server
	Attemps map[string]int
}

type LoginAttempt struct {
	Login    string
	Password string
}

func New() *authServer {
	cs := &authServer{}
	cs.Setup("Auth server", 4, 4)
	cs.Attemps = make(map[string]int)
	cs.RegisterInternalEndpoint(cs.login, 1)

	return cs
}

func (s *authServer) login(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("email", "password")
	if err != nil {
		answer.Data = err.Error()
	}
	attempt := &LoginAttempt{request.Data["email"], request.Data["password"]}
	key, err := s.NewAttempt(attempt)
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	answer.Ok = true
	answer.Data = key
	return answer

}

func (s *authServer) registerHandler(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("login", "email", "password")
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	err = s.createAccount(request.Data["login"], request.Data["email"], request.Data["password"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	answer.Data = "Account created!"
	answer.Ok = true
	return answer

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
func (s *authServer) checkCredentials(attempt *LoginAttempt) (models.User, error) {
	var err error
	var user models.User
	if len(attempt.Password) < 8 || len(attempt.Login) < 8 {
		err = errors.New(emptyFields)
	} else if user, err := storage.GetUserInfo(s.GetConn(), attempt.Login); err != nil {
		err = errors.New(wrongAccountInfo)
	} else if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(attempt.Password)); err != nil {
		err = errors.New(wrongAccountInfo)
	}
	return user, err

}
func (s *authServer) createAccount(login string, email string, password string) error {
	if len(login) < 6 || len(email) < 8 || len(password) < 8 {
		return errors.New(shortCredentials)
	} else if passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10); err != nil {
		return errors.New(InternalError)
	} else if err = storage.CreateUser(s.GetConn(), login, string(passwordHash), email); err != nil {
		return errors.New(accountExists + err.Error())
	}
	return nil

}
