package cityserver

import (
	"context"
	"github.com/joaopedrosgs/OpenLoU/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
)

func (cs *cityServer) upgradeConstructionAction(cityX int, cityY int, x int, y int) error {
	queueItem := models.Queue{ConstructionX: x, ConstructionY: y, CityX: cityX, CityY: cityY}
	queueItem.Insert(context.Background(), cs.GetConn(), boil.Infer())
	return queueItem.Insert(context.Background(), cs.GetConn(), boil.Infer())

}

func (cs *cityServer) newConstructionAction(cityX int, cityY int, x int, y int, cType int) error {

	city, err := models.FindCity(context.Background(), cs.GetConn(), cityX, cityY, "QueueTime")

	if err != nil {
		return err
	}

	construction := models.Construction{X: x, Y: y, CityX: cityX, CityY: cityY, Type: cType, Level: 0}
	err = construction.Insert(context.Background(), cs.GetConn(), boil.Infer())

	if err != nil {
		return err

	}

	if city.QueueTime.Before(time.Now()) {
		city.QueueTime = time.Now()
		city.Update(context.Background(), cs.GetConn(), boil.Infer())
	}
	queueItem := models.Queue{ConstructionX: x, ConstructionY: y, CityX: cityX, CityY: cityY, Action: 1, Completion: city.QueueTime.Add(time.Second * 10)}
	return queueItem.Insert(context.Background(), cs.GetConn(), boil.Infer())
}

func (cs *cityServer) getConstructionsAction(cityX int, cityY int) (constructions models.ConstructionSlice, err error) {

	constructions, err = models.Constructions(
		qm.Where("city_x=? AND city_y=?", cityX, cityY)).All(context.Background(), cs.GetConn())
	if err != nil {
		return
	}
	if len(constructions) > 0 {
		return
	}
	err = cs.newConstructionAction(cityX, cityY, 10, 10, 0)
	constructions = append(constructions, &models.Construction{X: 10, Y: 10, CityX: cityX, CityY: cityY, Type: 0, Level: 1})

	return
}

func (cs *cityServer) getUpgradesAction(cityX int, cityY int) (queueItems models.QueueSlice, err error) {

	queueItems, err = models.Queues(
		qm.Where("city_x=? AND city_y=?", cityX, cityY)).All(context.Background(), cs.GetConn())

	return
}
