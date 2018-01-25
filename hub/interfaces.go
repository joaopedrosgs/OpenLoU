package hub

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
)

type IWorker interface {
	GetName() string
	GetInChan() *chan *communication.Request
	SetOutChan(*chan *communication.Answer)
	GetCode() int
}
