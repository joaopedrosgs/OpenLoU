package cityserver

import (
	"context"
	"time"

	"github.com/joaopedrosgs/OpenLoU/server"
	"github.com/joaopedrosgs/OpenLoU/storage"
)

type cityServer struct {
	server.Server
}

func (cs *cityServer) UpgradeChecker() {
	for {
		upgrades, err := storage.GetAllUpgrades(cs.GetConn())
		if err != nil {
			cs.LogContext.Error(err.Error())
			time.Sleep(10 * time.Second)
			continue
		}
		batch := cs.GetConn().BeginBatch()
		for _, upgrade := range upgrades {
			upgrade.EnqueueCompletion(batch)
		}
		err = batch.Send(context.Background(), nil)
		if err != nil {
			cs.LogContext.Error(err.Error())
			time.Sleep(10 * time.Second)
			continue
		}
		time.Sleep(5 * time.Second)
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
