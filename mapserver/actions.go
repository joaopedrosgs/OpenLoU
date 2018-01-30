package mapserver

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/session"
)

func (ms *mapserver) createCity(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {

}
func (ms *mapserver) getCitiesFromUser(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {

	userID := session.GetUserID(request.Key)
	cities, err := database.GetUserCities(userID)
	if err != nil {
		answer.Data = err
		*out <- answer
		return
	}
	answer.Data = cities
	answer.Ok = true
	*out <- answer

}

func (ms *mapserver) getCities(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {
	err := request.ValidadeFields("X", "Y", "Range", "Continent")
	if err != nil {
		answer.Data = err.Error()
		*out <- answer
		return
	}
	if request.Data["X"] > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "Bad X value"
		*out <- answer
		return
	}
	if request.Data["Y"] > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "Bad Y value"
		*out <- answer
		return
	}

	if request.Data["Range"] <= 0 || request.Data["Range"] > 20 {
		request.Data["Range"] = 20
	}

	if err != nil || request.Data["Continent"] < 0 || request.Data["Continent"] > configuration.GetSingleton().Parameters.General.WorldSize {
		answer.Data = "Bad Continent value"
		*out <- answer
		return
	}
	cities, err := database.GetCitiesInRange(request.Data["X"], request.Data["Y"], request.Data["Range"], request.Data["Continent"])
	if err != nil {
		ms.LogContext.WithField("When", "Accessing database").Error(err.Error())
		answer.Data = "Internal error"
		*out <- answer
		return
	}
	answer.Ok = true
	answer.Data = cities
	*out <- answer

}
