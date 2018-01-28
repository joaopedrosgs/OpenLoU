package cityserver

import (
	"errors"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/database"
	"strconv"
)

func (cs *cityserver) upgradeConstruction(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {
	defer func() { *out <- answer }()
	cityID, err := strconv.ParseUint(request.Data["CityID"], 10, 32)
	if err != nil {
		answer.Data = BadCityID
		return
	}
	x, err := strconv.ParseUint(request.Data["X"], 10, 32)
	if err != nil {
		answer.Data = BadXValue
		return
	}
	y, err := strconv.ParseUint(request.Data["Y"], 10, 32)
	if err != nil {
		answer.Data = BadYValue
		return
	}
	err = database.CreateUpgrade(uint(cityID), uint(x), uint(y))
	if err != nil {
		answer.Data = err.Error()
		return
	}
	answer.Data = Success
	answer.Ok = true

}

func (cs *cityserver) newConstruction(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {
	defer func() { *out <- answer }()
	cityID, err := strconv.ParseUint(request.Data["CityID"], 10, 32)
	if err != nil {
		answer.Data = BadXValue
		return
	}
	x, err := strconv.ParseUint(request.Data["X"], 10, 32)
	if err != nil {
		answer.Data = BadXValue
		return
	}
	y, err := strconv.ParseUint(request.Data["Y"], 10, 32)
	if err != nil {
		answer.Data = BadYValue
		return
	}
	constructionType, err := strconv.ParseUint(request.Data["Type"], 10, 32)
	if err != nil {
		answer.Data = BadConstructionType
		return
	}
	_, err = database.GetConstruction(uint(cityID), uint(x), uint(y))
	if err != nil {
		err = database.CreateConstruction(uint(cityID), uint(x), uint(y), uint(constructionType), 0)
		if err != nil {
			answer.Data = errors.New(FailedNewConstruction)
			return
		}
		err = database.CreateUpgrade(uint(cityID), uint(x), uint(y))
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
	cityID, err := strconv.ParseUint(request.Data["CityID"], 10, 32)
	if err != nil {
		answer.Data = BadCityID
		return
	}
	cities, err := database.GetAllConstruction(uint(cityID))
	if err != nil {
		answer.Data = errors.New(FailedGetConstructions)
		return
	}
	answer.Ok = true
	answer.Data = cities

}
