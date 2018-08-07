package hub

import (
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/communication"
)

type IServer interface {
	GetName() string
	GetInChan() *chan *communication.Request
	SetOutChan(*chan *communication.Answer)
	GetCode() int
	SetConn(conn *pgx.Conn)
	StartListening()
}
