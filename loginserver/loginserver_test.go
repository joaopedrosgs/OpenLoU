package loginserver_test

import (
	"testing"

	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/loginserver"
)

var config configuration.Config
var attemptsArray = []struct {
	in  loginserver.LoginAttempt
	out bool
}{{loginserver.LoginAttempt{"", ""}, false},
	{loginserver.LoginAttempt{"wrong", "12345678"}, false},
	{loginserver.LoginAttempt{"test", "wrong"}, false},
	{loginserver.LoginAttempt{"test", ""}, false},
	{loginserver.LoginAttempt{"", "wrong"}, false},
	{loginserver.LoginAttempt{"test", "12345678"}, true}}

func TestLoginServer_NewAttempt(t *testing.T) {
	config.Load("../default.json")
	ls, _ := loginserver.New(&config)
	answer := loginserver.Answer{}
	_, err := ls.NewUser("test", "12345678", "testing@purpose.com")
	if err != nil {
		t.Error(err.Error())
	}
	for _, attempt := range attemptsArray {
		answer = ls.NewAttempt(attempt.in)
		if answer.Auth != attempt.out {
			t.Error("Unexpected result: (%s) != (%s)", answer.Auth, attempt.out)
		}
	}
	ls.DeleteUserByLogin("test")

}
func TestLoginServer_SessionExists(t *testing.T) {
	ls, _ := loginserver.New(&config)
	user, err := ls.NewUser("test", "12345678", "testing@purpose.com")
	if err != nil {
		t.Error(err.Error())
	}
	loginAttempt := loginserver.LoginAttempt{"127.0.0.1", user.Login, "12345678"}
	a := ls.NewAttempt(loginAttempt)
	if a.Auth {
		err := ls.SessionExists(loginserver.Session{UID: user.Id, Key: a.Key, Ip: "127.0.0.1"})
		if err != nil {
			t.Error("Expected session to exist, error: " + err.Error())
		}
	} else {
		t.Error("Expected to login normally")
	}
	ls.DeleteUserByLogin("test")

}
