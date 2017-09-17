package loginserver_test

import (
	"github.com/joaopedrosgs/OpenLoU/loginserver"
	"testing"
)

func TestCheckCredentials(t *testing.T) {
	attempt := loginserver.LoginAttempt{}
	err := loginserver.CheckCredentials(attempt)
	if err == nil {
		t.Error("Expected 'Empty fields' error")
	}
	attempt = loginserver.LoginAttempt{"127.0.0.1", "login", "pass"}
	err = loginserver.CheckCredentials(attempt)
	if err != nil {
		t.Error("Expected to login normally")
	}
}

func TestLoginServer_NewAttempt(t *testing.T) {
	attempt := loginserver.LoginAttempt{"127.0.0.1", "", ""}
	s := loginserver.LoginServer{}
	s.Sessions = make(map[string]loginserver.Session)
	answer := s.NewAttempt(attempt)
	if answer.Auth {
		t.Error("Expected empty Answer struct")
	}

}
func TestLoginServer_SessionExists(t *testing.T) {
	s := loginserver.LoginServer{}
	s.Sessions = make(map[string]loginserver.Session)
	loginAttempt := loginserver.LoginAttempt{"127.0.0.1", "login", "pass"}
	a := s.NewAttempt(loginAttempt)
	if a.Auth {
		err := s.SessionExists(loginserver.Auth{"login", a.Key, "127.0.0.1"})
		if err != nil {
			t.Error("Expected session to exist, error: " + err.Error())
		}
	} else {
		t.Error("Expected to login normally")
	}
}
