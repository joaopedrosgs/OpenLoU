package cityserver

import (
	core "github.com/joaopedrosgs/loucore/pkg"
	"github.com/joaopedrosgs/openlou/communication"
	"strconv"
)

func (cs *cityServer) upgradeConstruction(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("ConstructionID", "CityID")
	if err != nil {
		answer.Data = err.Error()
		return answer
	}

	cityId, err := strconv.Atoi(request.Data["CityID"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	constructionID, err := strconv.Atoi(request.Data["ConstructionID"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}

	qi, err := core.CreateQueueItem(request.GetSession().User.ID,cityId, constructionID,1)

	if err != nil {
		answer.Data = err.Error()
	} else {
		answer.Data = qi
		answer.Result = true
	}
	return answer

}

func (cs *cityServer) newConstruction(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("CityID", "X", "Y", "Type")
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
	cityId, err := strconv.Atoi(request.Data["CityID"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cType, err := strconv.Atoi(request.Data["Type"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}

	construction, err := core.CreateConstruction(cityId, x, y, cType)
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
	err := request.FieldsExist("CityID")
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cityId, err := strconv.Atoi(request.Data["CityID"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	constructions, err := core.GetStructuresFromCity(cityId)
	if err != nil {
		answer.Data = err.Error()
		return answer
	}

	answer.Result = true
	answer.Data = constructions

	return answer
}

func (cs *cityServer) getUpgrades(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	err := request.FieldsExist("CityID")
	if err != nil {
		answer.Data = err.Error()
		return answer
	}
	cityId, err := strconv.Atoi(request.Data["CityID"])
	if err != nil {
		answer.Data = err.Error()
		return answer
	}

	upgrades, err := core.GetCityQueue(cityId)
	if err != nil {
		answer.Data = err.Error()
	} else {
		answer.Result = true
		answer.Data = upgrades
	}
	return answer
}
