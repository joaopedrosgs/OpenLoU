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

	city := models.City{
		X:         0,
		Y:         0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		UserName: null.String{
			String: user.Name,
			Valid:  true,
		},
	}
	city.Insert(context.Background(), cs.GetConn(), boil.Infer())
	cs.LogContext.Info("Admin city created!")

	cityHall := models.Construction{
		X:         10,
		Y:         10,
		CityX:     city.X,
		CityY:     city.Y,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Level:     1,
		Type:      0,
	}
	cityHall.Insert(context.Background(), cs.GetConn(), boil.Infer())
	cs.LogContext.Info("Admin City Hall created!")

	cityHallUpgrade := models.Upgrade{
		ConstructionX: cityHall.X,
		ConstructionY: cityHall.Y,
		CityX:         cityHall.CityX,
		CityY:         cityHall.CityY,
		IndexAtQueue:  0,
		Duration:      10,
		Start:         time.Now(),
	}
	cityHallUpgrade.Insert(context.Background(), cs.GetConn(), boil.Infer())
	cs.LogContext.Info("Admin City Hall upgrade created!")

}
