package cityserver

import (
	"context"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
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
	upgrade := models.Upgrade{ConstructionX: x, ConstructionY: y, CityX: cityX, CityY: cityY, Duration: 10}
	upgrade.Insert(context.Background(), cs.GetConn(), boil.Infer())
	err = upgrade.Insert(context.Background(), cs.GetConn(), boil.Infer())
	if err != nil {
		answer.Data = err.Error()
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
	construction := models.Construction{X: x, Y: y, CityX: cityX, CityY: cityY, Type: cType}
	err = construction.Insert(context.Background(), cs.GetConn(), boil.Infer())
	if err != nil {
		answer.Data = err.Error()
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
	constructions, err := models.Constructions(
		qm.Where("city_x=? AND city_y=?", cityX, cityY)).All(context.Background(), cs.GetConn())
	if err != nil {
		answer.Data = err.Error()
	} else {
		answer.Result = true
		answer.Data = constructions
	}

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
	upgrades, err := models.Upgrades(
		qm.Where("city_x=? AND city_y=?", cityX, cityY)).All(context.Background(), cs.GetConn())
	if err != nil {
		answer.Data = err.Error()
	} else {
		answer.Result = true
		answer.Data = upgrades
	}
	return answer
}
