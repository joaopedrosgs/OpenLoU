package cityserver

import (
	"github.com/joaopedrosgs/OpenLoU/worker"
)

type cityserver struct {
	worker.Worker
}

func New() *cityserver {
	cs := &cityserver{}
	cs.Setup("City server", 2)
	return cs
}
