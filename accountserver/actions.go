package accountserver

import (
	"context"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (cs *accountServer) GetUserInfo(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()

	answer.Data = ""
	answer.Result = true
	return answer

}

func (cs *accountServer) CreateAdminAccount() {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin"), 10)
	user := models.User{
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "admin",
		Email:     "admin@admin",
		Password:  string(password),
		Gold:      0,
		Diamonds:  0,
		Darkwood:  0,
		Runestone: 0,
		Veritium:  0,
		Trueseed:  0,
		Rank:      0,
		AllianceName: null.String{
			String: "",
			Valid:  false,
		},
		AllianceRank: null.String{
			String: "",
			Valid:  false,
		},
	}
	user.Insert(context.Background(), cs.GetConn(), boil.Infer())
	cs.LogContext.Info("Admin account created!")

}
