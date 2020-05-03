package cityserver

import (
	"context"
	"github.com/joaopedrosgs/OpenLoU/server"
	"github.com/joaopedrosgs/openlou/ent/queue"
	"time"
)

type cityServer struct {
	server.Server
}

func (cs *cityServer) UpgradeChecker() {
	cs.LogContext.Info("Starting to check upgrades")
	for {

		time.Sleep(time.Second * 4)
		time := time.Now()

		queues, err := cs.GetClient().Queue.Query().WithConstruction().Where(queue.CompletionLTE(time)).All(context.Background())
		if err != nil {
			cs.LogContext.Error(err.Error())
			continue
		}
		_, err = cs.GetClient().Queue.Delete().Where(queue.CompletionLTE(time)).Exec(context.Background())

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
			if err != nil {
				cs.LogContext.Error(err.Error())
				continue
			}
			queue.Edges.Construction.Update().AddLevel(queue.Action).Save(context.Background())
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
