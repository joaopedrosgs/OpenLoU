package cityserver

import (
	"errors"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/entities"
)

func (cs *cityServer) upgradeConstruction(request *communication.Request, answer *communication.Answer) *communication.Answer {
	if err := request.FieldsExist("CityX", "CityY", "X", "Y"); err != nil {
		answer.Data = err.Error()
	} else if err := database.CreateUpgrade(&entities.Construction{
		X:     request.Data["X"],
		Y:     request.Data["Y"],
		CityX: request.Data["CityX"],
		CityY: request.Data["CityY"]}); err != nil {
		answer.Data = err.Error()
	} else {
		answer.Data = Success
		answer.Ok = true
	}
	return answer

}

func (cs *cityServer) newConstruction(request *communication.Request, answer *communication.Answer) *communication.Answer {

	if err := request.FieldsExist("CityX", "CityY", "X", "Y", "Type"); err != nil {
		answer.Data = err.Error()
	} else if construction, err := request.ToConstruction(); err != nil {
		answer.Data = err.Error()
	} else if _, err := database.GetConstruction(construction.CityX, construction.CityY, construction.X, construction.Y); err == nil {
		answer.Data = errors.New(TileInUse)
	} else if err = database.CreateConstruction(construction); err != nil {
		answer.Data = errors.New(FailedNewConstruction)
	} else if err = database.CreateUpgrade(construction); err != nil {
		answer.Data = errors.New(FailedNewConstructionUpgrade)
	} else {
		answer.Data = Success
		answer.Ok = true
	}
	return answer
}

func (cs *cityServer) getConstructions(request *communication.Request, answer *communication.Answer) *communication.Answer {
	err := request.FieldsExist("CityX", "CityY")
	if err != nil {
		answer.Data = err.Error()
	} else if cities, err := database.GetAllConstruction(request.Data["CityX"], request.Data["CityY"]); err != nil && cities != nil {
		answer.Data = errors.New(FailedGetConstructions)
	} else {
		answer.Ok = true
		answer.Data = cities
	}
	return answer
}
func (cs *cityServer) getUpgrades(request *communication.Request, answer *communication.Answer) *communication.Answer {
	err := request.FieldsExist("CityX", "CityY")
	if err != nil {
		answer.Data = err.Error()
	} else if upgrades, err := database.GetUpgradesFromCity(request.Data["CityX"], request.Data["CityY"]); err != nil && upgrades != nil {
		answer.Data = errors.New("Failed to get upgrads")
	} else {
		answer.Ok = true
		answer.Data = upgrades
	}
	return answer
}
