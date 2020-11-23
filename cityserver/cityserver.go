package cityserver

import (
	"github.com/joaopedrosgs/openlou/server"
	"time"
)

type cityServer struct {
	server.Server
}

func (cs *cityServer) UpgradeChecker() {
	cs.LogContext.Info("Starting to check upgrades")
	for {

		time.Sleep(time.Second * 4)


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
