package mapserver

import (
	"context"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"strconv"

	"github.com/joaopedrosgs/OpenLoU/models"
)

func (ms *mapserver) createCity(request *communication.Request) *communication.Answer {
	return request.ToAnswer()
}
func (ms *mapserver) getCitiesFromUser(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()

	cities, err := models.Cities(qm.Where("user_name=?", request.GetSession().User.Name)).All(context.Background(), ms.GetConn())
	if err != nil {
		ms.LogContext.WithField("When", "Accessing database to get cities from user").Error(err.Error())
		answer.Data = "Internal error"
		return answer

	}
	answer.Data = cities
	answer.Result = true

	return answer

}

func (ms *mapserver) getCities(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	if err := request.FieldsExist("X", "Y", "Range"); err != nil {
		answer.Data = err.Error()
		return answer
	}
	x, _ := strconv.Atoi(request.Data["X"])
	y, _ := strconv.Atoi(request.Data["Y"])
	rang, _ := strconv.Atoi(request.Data["Range"])

	if x < 0 || x > 600 {
		answer.Data = "X is bigger than allowed"
		return answer
	}
	if y < 0 || y > 600 {
		answer.Data = "Y is bigger than allowed"
		return answer

	}
	if rang < 1 || rang > 20 {
		answer.Data = "Range is bigger/smaller than allowed"
		return answer

	}
	cities, err := models.Cities(qm.Where(
		"x>? AND x<? AND y>? AND y<?", x-rang, x+rang, y-rang, y+rang)).All(context.Background(), ms.GetConn())
	if err != nil {
		ms.LogContext.WithField("When", "Accessing database to get cities in range").Error(err.Error())
		answer.Data = "Internal error"
		return answer

	}

	answer.Result = true
	answer.Data = cities

	return answer

}
