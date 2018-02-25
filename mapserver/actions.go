package mapserver

import (
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/hub"
	"github.com/joaopedrosgs/OpenLoU/session"
)

func (ms *mapserver) createCity(request *hub.RequestWithCallback) {
	return
}
func (ms *mapserver) getCitiesFromUser(request *hub.RequestWithCallback) {
	answer := request.Request.ToAnswer()
	userID := session.GetUserID(request.Request.Key)
	cities, err := database.GetUserCities(userID)
	if err != nil {
		answer.Data = err

	} else {
		answer.Data = cities
		answer.Ok = true
	}
	request.Callback(answer)
}

func (ms *mapserver) getCities(request *hub.RequestWithCallback) {
	answer := request.Request.ToAnswer()
	if err := request.Request.FieldsExist("X", "Y", "Range", "Continent"); err != nil {
		answer.Data = err.Error()
	}
	if request.Request.Data["X"] < 0 || request.Request.Data["X"] > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "X is bigger than allowed"
		return
	}
	if request.Request.Data["Y"] < 0 || request.Request.Data["Y"] > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "Y is bigger than allowed"
		return
	}
	if request.Request.Data["Range"] <= 0 || request.Request.Data["Range"] > 20 {
		answer.Data = "Range is bigger/smaller than allowed"
		return
	}
	if request.Request.Data["Continent"] < 0 || request.Request.Data["Continent"] > configuration.GetSingleton().Parameters.General.WorldSize {
		answer.Data = "This continent doenst exist"
		return
	}

	cities, err := database.GetCitiesInRange(request.Request.Data["X"], request.Request.Data["Y"], request.Request.Data["Range"], request.Request.Data["Continent"])
	if err != nil {
		ms.LogContext.WithField("When", "Accessing database").Error(err.Error())
		answer.Data = "Internal error"
	} else {
		answer.Ok = true
		answer.Data = cities
	}

	request.Callback(answer)
}
