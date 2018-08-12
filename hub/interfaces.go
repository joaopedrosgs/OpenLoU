package hub

import (
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
)

type IServer interface {
	GetName() string
	GetJobsChan() *chan *models.Request
	GetCode() int
	SetConn(conn *pgx.Conn)
	StartListening()
}
