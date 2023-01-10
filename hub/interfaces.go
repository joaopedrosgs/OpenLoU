package hub

import (
	"openlou/communication"
	"openlou/ent"
)

type IServer interface {
	GetName() string
	GetJobsChan() *chan *communication.Request
	GetCode() int
	SetClient(conn *ent.Client)
	StartListening()
	AfterSetup()
}
