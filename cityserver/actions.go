package cityserver

import (
	"errors"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/entities"
	"github.com/joaopedrosgs/OpenLoU/hub"
)

func (cs *cityserver) upgradeConstruction(request *hub.RequestWithCallback) {
	answer := request.Request.ToAnswer()
	defer request.Callback(answer)

	err := request.Request.FieldsExist("CityID ", "X", "Y")
	if err != nil {
		answer.Data = err.Error()
		return
	}
	err = database.CreateUpgrade(request.Request.Data["CityID"], request.Request.Data["X"], request.Request.Data["Y"])
	if err != nil {
		answer.Data = err.Error()
		return
	}
	answer.Data = Success
	answer.Ok = true
	return

}

func (cs *cityserver) newConstruction(request *hub.RequestWithCallback) {
	answer := request.Request.ToAnswer()
	defer request.Callback(answer)
	err := request.Request.FieldsExist("CityID", "X", "Y", "Type")
	if err != nil {
		answer.Data = err.Error()
		return
	}
	if request.Request.Data["X"] < 0 || request.Request.Data["X"] > 19 {
		answer.Data = "Bad X value"
		return
	}
	if request.Request.Data["Y"] < 0 || request.Request.Data["Y"] > 19 {
		answer.Data = "Bad Y value"
		return
	}
	_, ok := entities.RegisteredConstructions[request.Request.Data["Type"]]
	if !ok {
		answer.Data = "Bad Type value"
		return
	}
	_, err = database.GetConstruction(request.Request.Data["CityID"], request.Request.Data["X"], request.Request.Data["Y"]) // Check if construction exists
	if err == nil {
		answer.Data = errors.New(TileInUse)
		return
	}
	err = database.CreateConstruction(request.Request.Data["CityID"], request.Request.Data["X"], request.Request.Data["Y"], request.Request.Data["Type"], 0)
	if err != nil {
		answer.Data = errors.New(FailedNewConstruction)
		return
	}
	err = database.CreateUpgrade(request.Request.Data["CityID"], request.Request.Data["X"], request.Request.Data["Y"])
	if err != nil {
		answer.Data = errors.New(FailedNewConstructionUpgrade)
		return
	}
	answer.Data = Success
	answer.Ok = true
	return
}

func (cs *cityserver) getConstructions(request *hub.RequestWithCallback) {
	answer := request.Request.ToAnswer()
	defer request.Callback(answer)
	err := request.Request.FieldsExist("CityID")
	if err != nil {
		answer.Data = err.Error()
		return
	}
	cities, err := database.GetAllConstruction(request.Request.Data["CityID"])
	if err != nil {
		answer.Data = errors.New(FailedGetConstructions)
		return
	}
	answer.Ok = true
	answer.Data = cities
	return
}
