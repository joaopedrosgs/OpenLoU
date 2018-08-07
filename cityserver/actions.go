package cityserver

import (
	"errors"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/storage"
)

func (cs *cityServer) upgradeConstruction(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("CityX", "CityY", "X", "Y")
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	upgrade, err := request.ToUpgrade()
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	err = storage.CreateUpgrade(cs.GetConn(), upgrade)
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	answer.Data = Success
	answer.Ok = true
	return answer

}

func (cs *cityServer) newConstruction(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("CityX", "CityY", "X", "Y", "Type")
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	construction, err := request.ToConstruction()
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	err = storage.CreateConstruction(cs.GetConn(), construction)
	if err != nil {
		answer.Data = errors.New(FailedNewConstruction)
		return answer
	}

	answer.Data = Success
	answer.Ok = true
	return answer
}

func (cs *cityServer) getConstructions(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("CityX", "CityY")
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cityCoord, err := request.ToCityCoord()
	if err != nil {
		answer.Data = err
		return answer
	}
	cities, err := storage.GetAllConstructions(cs.GetConn(), *cityCoord)
	if err != nil {
		answer.Data = errors.New(FailedGetConstructions)
		return answer
	}
	answer.Ok = true
	answer.Data = cities

	return answer
}
func (cs *cityServer) getUpgrades(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("CityX", "CityY")
	if err != nil {
		answer.Data = err.Error()
	}
	cityCoord, err := request.ToCityCoord()
	if err != nil {
		answer.Data = err
		return answer
	}
	upgrades, err := storage.GetUpgradesFromCity(cs.GetConn(), *cityCoord)
	if err != nil {
		answer.Data = errors.New("Failed to get upgrades")
	}
	answer.Ok = true
	answer.Data = upgrades

	return answer
}
