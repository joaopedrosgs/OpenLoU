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

	if userName, ok := session.GetUserName(request.Key); !ok {
		answer.Data = "User not found"
	} else if cities, err := database.GetUserCities(userName); err != nil {
		answer.Data = err
	} else {
		answer.Data = cities
		answer.Ok = true
	}
	return answer

}

func (ms *mapserver) getCities(request *communication.Request, answer *communication.Answer) *communication.Answer {
	if err := request.FieldsExist("X", "Y", "Range"); err != nil {
		answer.Data = err.Error()
	} else if request.Data["X"] < 0 || request.Data["X"] > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "X is bigger than allowed"
	} else if request.Data["Y"] < 0 || request.Data["Y"] > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "Y is bigger than allowed"
	} else if request.Data["Range"] < 1 || request.Data["Range"] > 20 {
		answer.Data = "Range is bigger/smaller than allowed"
	} else if cities, err := database.GetCitiesInRange(request.Data["X"], request.Data["Y"], request.Data["Range"]); err != nil {
		ms.LogContext.WithField("When", "Accessing database").Error(err.Error())
		answer.Data = "Internal error"
	} else {
		answer.Ok = true
		answer.Data = cities
	}
	return answer

}
