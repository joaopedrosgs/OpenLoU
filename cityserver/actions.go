package cityserver

import (
	"errors"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/entities"
)

func (cs *cityserver) upgradeConstruction(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {
	defer func() { *out <- answer }()
	err := request.ValidadeFields("CityID", "X", "Y")
	if err != nil {
		answer.Data = err.Error()
		return
	}
	err = database.CreateUpgrade(request.Data["CityID"], request.Data["X"], request.Data["Y"])
	if err != nil {
		answer.Data = err.Error()
		return
	}
	answer.Data = Success
	answer.Ok = true

}

func (cs *cityserver) newConstruction(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {
	defer func() { *out <- answer }()
	err := request.ValidadeFields("CityID", "X", "Y", "Type")
	if err != nil {
		answer.Data = err.Error()
		return
	}
	if request.Data["X"] > 19 {
		answer.Data = "Bad X value"
		*out <- answer
		return
	}
	if request.Data["Y"] > 19 {
		answer.Data = "Bad Y value"
		*out <- answer
		return
	}
	_, ok := entities.RegisteredConstructions[request.Data["Type"]]
	if !ok {
		answer.Data = "Bad Type value"
		*out <- answer
		return
	}
	_, err = database.GetConstruction(request.Data["CityID"], request.Data["X"], request.Data["Y"]) // Check if construction exists
	if err != nil {
		err = database.CreateConstruction(request.Data["CityID"], request.Data["X"], request.Data["Y"], request.Data["Type"], 0)
		if err != nil {
			answer.Data = errors.New(FailedNewConstruction)
			return
		}
		err = database.CreateUpgrade(request.Data["CityID"], request.Data["X"], request.Data["Y"])
		if err != nil {
			answer.Data = errors.New(FailedNewConstructionUpgrade)
			return
		}
		answer.Data = Success
		answer.Ok = true
	} else {
		answer.Data = errors.New(TileInUse)
	}

}

func (cs *cityserver) getConstructions(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {
	defer func() { *out <- answer }()
	err := request.ValidadeFields("CityID")
	if err != nil {
		answer.Data = err.Error()
		return
	}
	cities, err := database.GetAllConstruction(request.Data["CityID"])
	if err != nil {
		answer.Data = errors.New(FailedGetConstructions)
		return
	}
	answer.Ok = true
	answer.Data = cities

}
