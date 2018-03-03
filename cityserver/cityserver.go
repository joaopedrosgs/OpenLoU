package cityserver

import (
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/worker"
	"time"
)

type cityServer struct {
	worker.Worker
}

func (cs *cityServer) UpgradeChecker() {
	for {
		upgrades, err := database.GetUpgrades()
		if err != nil {
			cs.LogContext.Error(err.Error())
			time.Sleep(10 * time.Second)
			continue
		}
		for _, upgrade := range *upgrades {
			err := database.CompleteUpgrade(upgrade)
			if err != nil {
				cs.LogContext.Error(err.Error())
				time.Sleep(10 * time.Second)
				continue
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func New() *cityServer {
	cs := &cityServer{}
	cs.Setup("City server", 3)
	cs.RegisterInternalEndpoint(cs.upgradeConstruction, 1)
	cs.RegisterInternalEndpoint(cs.newConstruction, 2)
	cs.RegisterInternalEndpoint(cs.getConstructions, 3)
	cs.RegisterInternalEndpoint(cs.getUpgrades, 4)

	go cs.UpgradeChecker()
	return cs
}
