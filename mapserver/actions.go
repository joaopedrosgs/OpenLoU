package mapserver

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/session"
	"strconv"
)

func (ms *mapserver) createCity(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {

}
func (ms *mapserver) getCitiesFromUser(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {
	defer func() { *out <- answer }()
	userID := session.GetUserID(request.Key)
	cities, err := database.GetUserCities(userID)
	if err != nil {
		answer.Data = err
		return
	}
	answer.Data, answer.Ok = cities, true

}

func (ms *mapserver) getCities(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {
	x, err := strconv.Atoi(request.Data["X"])
	if err != nil || x < 0 || x > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "Bad X value"
		return
	}
	y, err := strconv.Atoi(request.Data["Y"])
	if err != nil || y < 0 || y > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "Bad Y value"
		return

	}
	distance, err := strconv.Atoi(request.Data["Range"])
	if err != nil || distance <= 0 || distance > 10 {
		answer.Data = "Bad Range value"
		return
	}

	continent, err := strconv.Atoi(request.Data["Continent"])
	if err != nil || continent < 0 || continent > 50 {
		answer.Data = "Bad Continent value"
		return
	}
	cities, err := database.GetCitiesInRange(x, y, distance, continent)
	if err != nil {
		answer.Data = err.Error()
		return
	}
	answer.Ok = true
	answer.Data = cities

}
