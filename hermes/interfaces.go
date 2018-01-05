package hermes

import "github.com/joaopedrosgs/OpenLoU/communication"

type ISessionBackend interface {
	NewSession(id int, key string)

	SessionExists(key string) bool

	DeleteSession(key string)

	NewTry(key string)
}

type IWorker interface {
	GetInChan() *chan *communication.Request
	SetOutChan(*chan *communication.Answer)
	GetCode() int
}
