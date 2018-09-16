package hub

import (
	"database/sql"
	"github.com/joaopedrosgs/OpenLoU/communication"
)

type IServer interface {
	GetName() string
	GetJobsChan() *chan *communication.Request
	GetCode() int
	SetConn(conn *sql.DB)
	StartListening()
	AfterSetup()
}
