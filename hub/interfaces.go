package hub

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/openlou/ent"
)

type IServer interface {
	GetName() string
	GetJobsChan() *chan *communication.Request
	GetCode() int
	SetClient(conn *ent.Client)
	StartListening()
	AfterSetup()
}
