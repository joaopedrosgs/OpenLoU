package mapserver

import (
	"github.com/joaopedrosgs/OpenLoU/worker"
)

type mapserver struct {
	worker.Worker
}

func New() *mapserver {
	ms := &mapserver{}
	ms.Setup("Map Server", 2)
	ms.RegisterInternalEndpoint(ms.createCity, 1)
	ms.RegisterInternalEndpoint(ms.getCities, 2)
	ms.RegisterInternalEndpoint(ms.getCitiesFromUser, 3)

	return ms
}
