package loginserver_test

import (
	"github.com/joaopedrosgs/OpenLoU/loginserver"
	"testing"
)

var attemptsArray = []struct {
	in  loginserver.LoginAttempt
	out bool
}{{loginserver.LoginAttempt{"127.0.0.1", "", ""}, false},
	{loginserver.LoginAttempt{"127.0.0.1", "wrong", "12345678"}, false},
	{loginserver.LoginAttempt{"127.0.0.1", "test", "wrong"}, false},
	{loginserver.LoginAttempt{"127.0.0.1", "test", ""}, false},
	{loginserver.LoginAttempt{"127.0.0.1", "", "wrong"}, false},
	{loginserver.LoginAttempt{"127.0.0.1", "test", "12345678"}, true}}

func TestLoginServer_NewAttempt(t *testing.T) {

	ls := loginserver.New(10)
	answer := loginserver.Answer{}
	ls.NewUser("test", "12345678", "testing@purpose.com")

	for _, ele := range attemptsArray {
		answer = ls.NewAttempt(ele.in)
		if answer.Auth != ele.out {
			t.Error("Unexpected result: (%s) != (%s)", answer.Auth, ele.out)
		}
	}
	ls.DeleteUserByLogin("test")

}
func TestLoginServer_SessionExists(t *testing.T) {

	ls := loginserver.New(10)
	ls.NewUser("test", "12345678", "testing@purpose.com")
	loginAttempt := loginserver.LoginAttempt{"127.0.0.1", "test", "12345678"}
	a := ls.NewAttempt(loginAttempt)
	if a.Auth {
		err := ls.SessionExists(loginserver.Session{Login: "login", Key: a.Key, Ip: "127.0.0.1"})
		if err != nil {
			t.Error("Expected session to exist, error: " + err.Error())
		}
	} else {
		t.Error("Expected to login normally")
	}
	ls.DeleteUserByLogin("test")

}
