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
	cs.RegisterInternalEndpoint(cs.GetUserInfo, 1)

	return cs
}
