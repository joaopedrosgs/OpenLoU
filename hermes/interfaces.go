package hermes

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/session"
	"net"
)

type ISessionBackend interface {
	NewSession(id int, key string)
	SessionExists(key string) bool
	DeleteSession(key string)
	NewTry(key string)
	GetUserId(key string) int
	GetUserConnByKey(key string) net.Conn
	GetUserConnByID(id int) net.Conn
	SetConn(key string, conn net.Conn)
	GetSession(key string) *session.Session
}

type IWorker interface {
	GetInChan() *chan *communication.Request
	SetOutChan(*chan *communication.Answer)
	GetCode() int
}
