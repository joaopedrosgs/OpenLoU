package cityserver

import (
	"errors"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/entities"
)

func (cs *cityserver) upgradeConstruction(request *communication.Request, answer *communication.Answer) *communication.Answer {
	if err := request.FieldsExist("CityID ", "X", "Y"); err != nil {
		answer.Data = err.Error()
	} else if err = database.CreateUpgrade(request.Data["CityID"], request.Data["X"], request.Data["Y"]); err != nil {
		answer.Data = err.Error()
	} else {
		answer.Data = Success
		answer.Ok = true
	}
	return answer

}

func (cs *cityserver) newConstruction(request *communication.Request, answer *communication.Answer) *communication.Answer {
	if err := request.FieldsExist("CityID", "X", "Y", "Type"); err != nil {
		answer.Data = err.Error()
	} else if request.Data["X"] < 0 || request.Data["X"] > 19 {
		answer.Data = BadXValue
	} else if request.Data["Y"] < 0 || request.Data["Y"] > 19 {
		answer.Data = BadYValue
	} else if _, ok := entities.RegisteredConstructions[request.Data["Type"]]; !ok {
		answer.Data = BadConstructionType
	} else if _, err = database.GetConstruction(request.Data["CityID"], request.Data["X"], request.Data["Y"]); err != nil {
		answer.Data = errors.New(TileInUse)
	} else if err = database.CreateConstruction(request.Data["CityID"], request.Data["X"], request.Data["Y"], request.Data["Type"], 0); err != nil {
		answer.Data = errors.New(FailedNewConstruction)
	} else if err = database.CreateUpgrade(request.Data["CityID"], request.Data["X"], request.Data["Y"]); err != nil {
		answer.Data = errors.New(FailedNewConstructionUpgrade)
	} else {
		answer.Data = Success
		answer.Ok = true
	}
	return answer
}

func (cs *cityserver) getConstructions(request *communication.Request, answer *communication.Answer) *communication.Answer {
	err := request.FieldsExist("CityID")
	if err != nil {
		answer.Data = err.Error()
	} else if cities, err := database.GetAllConstruction(request.Data["CityID"]); err != nil {
		answer.Data = errors.New(FailedGetConstructions)
	} else {
		answer.Ok = true
		answer.Data = cities
	}
	return answer
}
