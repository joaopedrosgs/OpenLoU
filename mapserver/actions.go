package mapserver

import (
	"strconv"

	"github.com/joaopedrosgs/OpenLoU/models"
	"github.com/joaopedrosgs/OpenLoU/storage"
)

func (ms *mapserver) createCity(request *models.Request) *models.Answer {
	return request.ToAnswer()
}
func (ms *mapserver) getCitiesFromUser(request *models.Request) *models.Answer {
	answer := request.ToAnswer()
	cities, err := storage.GetUserCities(ms.GetConn(), request.Session.User.Name)
	if err != nil {
		answer.Data = err
		return answer
	}
	answer.Data = cities
	answer.Ok = true

	return answer

}

func (ms *mapserver) getCities(request *models.Request) *models.Answer {
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
