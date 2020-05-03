package cityserver

import (
	"github.com/joaopedrosgs/openlou/communication"
	"strconv"
)

func (cs *cityServer) upgradeConstruction(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("CityX", "CityY", "X", "Y")
	if err != nil {
		answer.Data = err.Error()
		return answer
	}

	x, err := strconv.Atoi(request.Data["X"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	y, err := strconv.Atoi(request.Data["Y"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cityX, err := strconv.Atoi(request.Data["CityX"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cityY, err := strconv.Atoi(request.Data["CityY"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	_, err = cs.upgradeConstructionAction(cityX, cityY, x, y)

	if err != nil {
		answer.Data = err.Error()
	} else {
		answer.Result = true
	}
	return answer

}

func (cs *cityServer) newConstruction(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("CityX", "CityY", "X", "Y", "Type")
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	x, err := strconv.Atoi(request.Data["X"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	y, err := strconv.Atoi(request.Data["Y"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cityX, err := strconv.Atoi(request.Data["CityX"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cityY, err := strconv.Atoi(request.Data["CityY"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cType, err := strconv.Atoi(request.Data["Type"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}

	construction, err := cs.newConstructionAction(cityX, cityY, x, y, cType, request.GetSession())
	if err != nil {
		answer.Data = err.Error()
	} else {
		answer.Data = construction
		answer.Result = true
	}
	return answer
}

func (cs *cityServer) getConstructions(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("CityX", "CityY")
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cityX, err := strconv.Atoi(request.Data["CityX"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cityY, err := strconv.Atoi(request.Data["CityY"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	constructions, err := cs.getConstructionsAction(cityX, cityY, request.GetSession())
	if err != nil {
		answer.Data = err.Error()
		return answer
	}

	if len(constructions) == 0 {
		townHall, _ := cs.newConstructionAction(cityX, cityY, 10, 10, 0, request.GetSession())
		constructions = append(constructions, townHall)
	}

	answer.Result = true
	answer.Data = constructions

	return answer
}

func (cs *cityServer) getUpgrades(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("CityX", "CityY")
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cityX, err := strconv.Atoi(request.Data["CityX"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cityY, err := strconv.Atoi(request.Data["CityY"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	upgrades, err := cs.getUpgradesAction(cityX, cityY, request.GetSession())
	if err != nil {
		answer.Data = err.Error()
	} else {
		answer.Result = true
		answer.Data = upgrades
	}
	return answer
}
