package accountserver

import (
	"github.com/joaopedrosgs/openlou/server"
)

type accountServer struct {
	server.Server
}

func (cs *accountServer) AfterSetup() {
}

func New() *accountServer {
	cs := &accountServer{}
	cs.Setup("Account server", 1, 10)
	cs.RegisterInternalEndpoint(cs.GetUserInfo, 1)
	cs.RegisterInternalEndpoint(cs.CreateAccount, 2)
	cs.RegisterInternalEndpoint(cs.Login, 3)

	return cs
}
