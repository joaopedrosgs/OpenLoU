package loginserver_test

import (
	"github.com/joaopedrosgs/OpenLoU/loginserver"
	"strings"
	"testing"
)

func TestLoginServer_User(t *testing.T) {
	config.LoadDefault()
	l, _ := loginserver.New(&config)
	user1, err := l.NewUser("test", "1235678", "testing@purposes.com")
	if err != nil {
		t.Fatal(err.Error())
	}
	user1.Email = "purposes@testing.com"
	err = l.SaveUserChanges(user1)
	if err != nil {
		t.Fatal(err.Error())
	}
	user2, err := l.LoadUserByLogin("test")
	if err != nil {
		t.Fatal(err.Error())
	}
	if strings.Compare(user2.Email, user1.Email) != 0 {
		t.Error("Expected %s and %s to be equal", user2.Email, user1.Email)
	}
	err = l.DeleteUserByLogin("test")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = l.UserExists("test", "testing@purposes.com")
	if err == nil {
		t.Fatal("User should have been deleted, error: %s", err.Error())
	}

}
