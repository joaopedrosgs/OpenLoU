package accountserver

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"openlou/ent"
	"openlou/ent/user"
)

func (cs *accountServer) CreateAdminAccountAction() {
	cs.CreateAccountAction("admin@admin", "admin", "admin")

}

func (cs *accountServer) CreateAccountAction(email string, name string, password string) (*ent.User, error) {
	passCrypt, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return cs.GetClient().User.
		Create().
		SetName(name).
		SetEmail(email).
		SetPasswordHash(
			string(passCrypt)).
		Save(context.Background())

}
func (cs *accountServer) CheckAccountAction(email string, password string) (*ent.User, error) {
	account, err := cs.GetClient().User.
		Query().Where(user.EmailEQ(email)).Only(context.Background())
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(password)) != nil {
		return nil, errors.New("Wrong password")
	}
	return account, nil

}
