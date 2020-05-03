package mapserver

import (
	"github.com/joaopedrosgs/openlou/communication"
	"strconv"
)

func (ms *mapserver) createCity(request *communication.Request) *communication.Answer {
	return request.ToAnswer()
}
func (ms *mapserver) getCitiesFromUser(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()

	cities, err := ms.getCitiesFromUserAction(request.GetSession().User)
	if err != nil {
		answer.Data = "Failed to get cities from user:" + err.Error()
		return answer
	}
	if len(cities) == 0 {
		firstCity, err := ms.createCityAction(10, 10, request.GetSession().User)
		if err != nil {
			answer.Data = "Failed to get cities from user:" + err.Error()
			return answer
		}
		cities = append(cities, firstCity)
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

	cities, err := ms.getCitiesAction(x, y, rang)

	if err != nil {
		answer.Data = "Failed to get cities"
	} else {
		answer.Result = true
		answer.Data = cities
	}

	return answer

}
