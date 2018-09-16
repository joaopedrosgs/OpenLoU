package session

import (
	"github.com/joaopedrosgs/OpenLoU/models"
	"time"

	"github.com/gorilla/websocket"
)

type Session struct {
	LastAction time.Time
	tries      int
	Conn       *websocket.Conn
	User       *models.User
}

func (s *Session) NewTry() {
	s.tries++
}

func (s *Session) Close() {
	s.Conn.Close()
}
