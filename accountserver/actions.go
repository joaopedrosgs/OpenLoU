package accountserver

import (
	"context"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"golang.org/x/crypto/bcrypt"
)

func (cs *accountServer) GetUserInfo(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()

	answer.Data = ""
	answer.Result = true
	return answer

}

func (cs *accountServer) CreateAdminAccount() {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin"), 10)
	cs.GetClient().User.Create().SetName("admin").SetEmail("admin@admin").SetPasswordHash(string(password)).Save(context.Background())
	cs.LogContext.Info("Admin account created!")

}
