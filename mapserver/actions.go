package mapserver

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/session"
)

func (ms *mapserver) createCity(request *communication.Request, answer *communication.Answer) *communication.Answer {
	return answer
}
func (ms *mapserver) getCitiesFromUser(request *communication.Request, answer *communication.Answer) *communication.Answer {

	userID := session.GetUserID(request.Key)
	cities, err := database.GetUserCities(userID)
	if err != nil {
		answer.Data = err

	} else {
		answer.Data = cities
		answer.Ok = true
	}
	return answer

}

func (ms *mapserver) getCities(request *communication.Request, answer *communication.Answer) *communication.Answer {
	if err := request.FieldsExist("X", "Y", "Range", "Continent"); err != nil {
		answer.Data = err.Error()
	} else if request.Data["X"] < 0 || request.Data["X"] > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "X is bigger than allowed"
	} else if request.Data["Y"] < 0 || request.Data["Y"] > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "Y is bigger than allowed"
	} else if request.Data["Range"] <= 0 || request.Data["Range"] > 20 {
		answer.Data = "Range is bigger/smaller than allowed"
	} else if request.Data["Continent"] < 0 || request.Data["Continent"] > configuration.GetSingleton().Parameters.General.WorldSize {
		answer.Data = "This continent doenst exist"
	} else if cities, err := database.GetCitiesInRange(request.Data["X"], request.Data["Y"], request.Data["Range"], request.Data["Continent"]); err != nil {
		ms.LogContext.WithField("When", "Accessing database").Error(err.Error())
		answer.Data = "Internal error"
	} else {
		answer.Ok = true
		answer.Data = cities
	}
	return answer

}
