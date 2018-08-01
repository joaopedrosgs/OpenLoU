package accountserver

import (
	"github.com/joaopedrosgs/OpenLoU/server"
)

type accountServer struct {
	server.Server
}

func New() *accountServer {
	cs := &accountServer{}
	cs.Setup("City server", 1)
	return cs
}
