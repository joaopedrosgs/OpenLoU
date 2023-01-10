package mapserver

import (
	"openlou/server"
)

type mapserver struct {
	server.Server
}

func (ms *mapserver) AfterSetup() {
}

func New() *mapserver {
	ms := &mapserver{}
	ms.Setup("Map Server", 2, 4)
	ms.RegisterInternalEndpoint(ms.createCity, 1)
	ms.RegisterInternalEndpoint(ms.getCities, 2)
	ms.RegisterInternalEndpoint(ms.getCitiesFromUser, 3)
	return ms
}
