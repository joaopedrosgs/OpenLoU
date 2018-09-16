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

		upgrades, err := models.Upgrades(qm.Where("index_at_queue=0 AND (start+duration * interval '1 second')> CURRENT_TIMESTAMP")).All(context.Background(), tx)

		if err != nil {
			cs.LogContext.Error(err.Error())
			continue
		}
		cs.LogContext.Infof("%d upgrades found", len(upgrades))
		if len(upgrades) == 0 {
			continue
		}
		if err != nil {
			cs.LogContext.Error(err.Error())
			continue
		}
		for _, upgrade := range upgrades {

			construction, err := models.FindConstruction(context.Background(), cs.GetConn(), upgrade.CityX, upgrade.CityY, upgrade.ConstructionX, upgrade.ConstructionY)
			if err != nil {
				cs.LogContext.Error(err.Error())
				continue
			}
			construction.Level++
			construction.Update(context.Background(), tx, boil.Infer())
		}
		_, err = upgrades.DeleteAll(context.Background(), tx)
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
