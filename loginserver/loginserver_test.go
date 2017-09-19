package loginserver_test

import (
	"github.com/joaopedrosgs/OpenLoU/loginserver"
	"testing"
)

func TestCheckCredentials(t *testing.T) {
	attempt := loginserver.LoginAttempt{}
	ls := loginserver.New(10)
	err, _ := ls.CheckCredentials(attempt)
	if err == nil {
		t.Error("Expected 'Empty fields' error")
	}
	attempt = loginserver.LoginAttempt{"127.0.0.1", "pedro", "12345"}
	err, _ = ls.CheckCredentials(attempt)
	if err != nil {
		t.Error("Expected to login normally")
	}
}

func TestLoginServer_NewAttempt(t *testing.T) {

	ls := loginserver.New(10)
	attempt := loginserver.LoginAttempt{}
	answer := loginserver.Answer{}

	attempt = loginserver.LoginAttempt{"127.0.0.1", "", ""}
	answer = ls.NewAttempt(attempt)
	if answer.Auth {
		t.Error("Expected to not login")
	}
	attempt = loginserver.LoginAttempt{"127.0.0.1", "pedro", "12345"}
	answer = ls.NewAttempt(attempt)
	if !answer.Auth {
		t.Error("Expected to not login")
	}
	attempt = loginserver.LoginAttempt{"127.0.0.1", "wrong", "12345"}
	answer = ls.NewAttempt(attempt)
	if answer.Auth {
		t.Error("Expected to not login")
	}
	attempt = loginserver.LoginAttempt{"127.0.0.1", "pedro", "wrong"}
	answer = ls.NewAttempt(attempt)
	if answer.Auth {
		t.Error("Expected to not login")
	}

}
func TestLoginServer_SessionExists(t *testing.T) {

	ls := loginserver.New(10)
	loginAttempt := loginserver.LoginAttempt{"127.0.0.1", "pedro", "12345"}
	a := ls.NewAttempt(loginAttempt)
	if a.Auth {
		err := ls.SessionExists(loginserver.Session{Login: "login", Key: a.Key, Ip: "127.0.0.1"})
		if err != nil {
			t.Error("Expected session to exist, error: " + err.Error())
		}
	} else {
		t.Error("Expected to login normally")
	}
}
