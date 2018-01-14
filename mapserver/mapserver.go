package mapserver

import (
	"github.com/joaopedrosgs/OpenLoU/worker"
)

type mapserver struct {
	worker.Worker
}

func New() *mapserver {
	ms := &mapserver{}
	ms.Setup("Map Server", 1)
	ms.RegisterInternalEndpoint(ms.createCity, 101)
	ms.RegisterInternalEndpoint(ms.getCities, 102)

	return ms
}
