package hub

import (
	"github.com/joaopedrosgs/openlou/communication"
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
