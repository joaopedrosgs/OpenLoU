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
		answer.Data = "Bad CityID"
		return
	}
	x, err := strconv.ParseUint(request.Data["X"], 10, 32)
	if err != nil {
		answer.Data = "Bad X value"
		return
	}
	y, err := strconv.ParseUint(request.Data["Y"], 10, 32)
	if err != nil {
		answer.Data = "Bad Y value"
		return
	}
	err = database.CreateUpgrade(uint(cityID), uint(x), uint(y))
	if err != nil {
		answer.Data = err.Error()
		return
	}
	answer.Data = "success"
	answer.Ok = true

}

func (cs *cityserver) newConstruction(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {
	defer func() { *out <- answer }()
	cityID, err := strconv.ParseUint(request.Data["CityID"], 10, 32)
	if err != nil {
		answer.Data = "Bad CityID"
		return
	}
	x, err := strconv.ParseUint(request.Data["X"], 10, 32)
	if err != nil {
		answer.Data = "Bad X value"
		return
	}
	y, err := strconv.ParseUint(request.Data["Y"], 10, 32)
	if err != nil {
		answer.Data = "Bad Y value"
		return
	}
	constructionType, err := strconv.ParseUint(request.Data["Type"], 10, 32)
	if err != nil {
		answer.Data = "Bad construction type"
		return
	}
	_, err = database.GetConstruction(uint(cityID), uint(x), uint(y))
	if err != nil {
		err = database.CreateConstruction(uint(cityID), uint(x), uint(y), uint(constructionType), 0)
		if err != nil {
			answer.Data = errors.New("failed to create construction")
			return
		}
		err = database.CreateUpgrade(uint(cityID), uint(x), uint(y))
		if err != nil {
			answer.Data = errors.New("failed to create upgrade to the new construction")
			return
		}
		answer.Data = "success"
		answer.Ok = true
	} else {
		answer.Data = errors.New("tile already in use")
	}

}

func (cs *cityserver) getConstructions(request *communication.Request, answer *communication.Answer, out *chan *communication.Answer) {
	defer func() { *out <- answer }()
	cityID, err := strconv.ParseUint(request.Data["CityID"], 10, 32)
	if err != nil {
		answer.Data = "Bad CityID"
		return
	}
	cities, err := database.GetAllConstruction(uint(cityID))
	if err != nil {
		answer.Data = errors.New("failed to get constructions")
		return
	}
	answer.Data = cities

}
