package mapserver

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/database"
	"strconv"
)

func (ms *mapserver) createCity(requestData map[string]string, answer *communication.Answer) {

}
func (ms *mapserver) getCities(requestData map[string]string, answer *communication.Answer) {

	x, err := strconv.Atoi(requestData["X"])
	if err != nil || x < 0 || x > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "Bad X value"
		return
	}
	y, err := strconv.Atoi(requestData["Y"])
	if err != nil || y < 0 || y > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "Bad Y value"
		return
	}
	distance, err := strconv.Atoi(requestData["Range"])
	if err != nil || distance <= 0 || distance > 10 {
		answer.Data = "Bad Range value"
		return
	}

	continent, err := strconv.Atoi(requestData["Continent"])
	if err != nil || continent < 0 || continent > 50 {
		answer.Data = "Bad Continent value"
		return
	}
	answer.Data = database.GetCitiesInRange(x, y, distance, continent)

}
