package mapserver

import (
	"strconv"

	"github.com/joaopedrosgs/OpenLoU/models"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/session"
	"github.com/joaopedrosgs/OpenLoU/storage"
)

func (ms *mapserver) createCity(request *communication.Request) *communication.Answer {
	return request.ToAnswer()
}
func (ms *mapserver) getCitiesFromUser(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	userName, err := session.GetUserName(request.Key)
	if err != nil {
		answer.Data = "User not found: " + err.Error()
		return answer
	}
	cities, err := storage.GetUserCities(ms.GetConn(), userName)
	if err != nil {
		answer.Data = err
		return answer
	}
	answer.Data = cities
	answer.Ok = true

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
	cities, err := storage.GetAllCitiesInRange(ms.GetConn(), models.Coord{x, y}, rang)
	if err != nil {
		ms.LogContext.WithField("When", "Accessing database").Error(err.Error())
		answer.Data = "Internal error"
		return answer

	}

	answer.Ok = true
	answer.Data = cities

	return answer

}
