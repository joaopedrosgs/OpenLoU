package cityserver

import (
	"github.com/joaopedrosgs/OpenLoU/database"

	"github.com/joaopedrosgs/OpenLoU/worker"
	"time"
)

type cityserver struct {
	worker.Worker
}

func (cs *cityserver) UpgradeChecker() {
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

func New() *cityserver {
	cs := &cityserver{}
	cs.Setup("City server", 2)
	cs.RegisterInternalEndpoint(cs.upgradeConstruction, 1)
	cs.RegisterInternalEndpoint(cs.newConstruction, 2)
	cs.RegisterInternalEndpoint(cs.getConstructions, 3)


	go cs.UpgradeChecker()
	return cs
}
