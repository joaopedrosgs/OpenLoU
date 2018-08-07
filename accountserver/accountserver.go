package accountserver

import (
	"github.com/joaopedrosgs/OpenLoU/server"
)

type accountServer struct {
	server.Server
}

func New() *accountServer {
	cs := &accountServer{}
	cs.Setup("Account server", 1, 10)
	return cs
}
