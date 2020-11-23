package hub

import (
	"github.com/joaopedrosgs/openlou/communication"
)

type IServer interface {
	GetName() string
	GetJobsChan() *chan *communication.Request
	GetCode() int
	StartListening()
	AfterSetup()
}
