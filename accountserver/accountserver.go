package accountserver

import (
	"github.com/joaopedrosgs/OpenLoU/worker"
)

type accountServer struct {
	worker.Worker
}

func New() *accountServer {
	cs := &accountServer{}
	cs.Setup("City server", 1)
	return cs
}
