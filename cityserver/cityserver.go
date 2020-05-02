package cityserver

import (
	"context"
	"github.com/joaopedrosgs/OpenLoU/models"
	"github.com/joaopedrosgs/OpenLoU/server"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
)

type cityServer struct {
	server.Server
}

func (cs *cityServer) UpgradeChecker() {
	cs.LogContext.Info("Starting to check upgrades")
	for {
		time.Sleep(time.Second * 4)

		tx, err := cs.GetConn().Begin()

		queues, err := models.Queues(qm.Where("completion < CURRENT_TIMESTAMP")).All(context.Background(), tx)

		if err != nil {
			cs.LogContext.Error(err.Error())
			continue
		}
		cs.LogContext.Infof("%d queue item(s) found", len(queues))
		if len(queues) == 0 {
			continue
		}
		if err != nil {
			cs.LogContext.Error(err.Error())
			continue
		}
		for _, queue := range queues {

			construction, err := models.FindConstruction(context.Background(), cs.GetConn(), queue.CityX, queue.CityY, queue.ConstructionX, queue.ConstructionY)
			if err != nil {
				cs.LogContext.Error(err.Error())
				continue
			}
			cs.LogContext.Info("Date ", queue.Completion)
			construction.Level += queue.Action
			construction.Update(context.Background(), tx, boil.Infer())
		}
		_, err = queues.DeleteAll(context.Background(), tx)
		if err != nil {
			cs.LogContext.Error(err.Error())
			tx.Rollback()

		}
		err = tx.Commit()
		if err != nil {
			cs.LogContext.Error(err.Error())

		}

	}
}

func New() *cityServer {
	cs := &cityServer{}
	cs.Setup("City server", 3, 4)
	cs.RegisterInternalEndpoint(cs.upgradeConstruction, 1)
	cs.RegisterInternalEndpoint(cs.newConstruction, 2)
	cs.RegisterInternalEndpoint(cs.getConstructions, 3)
	cs.RegisterInternalEndpoint(cs.getUpgrades, 4)
	return cs
}
func (cs *cityServer) AfterSetup() {
	go cs.UpgradeChecker()
}
