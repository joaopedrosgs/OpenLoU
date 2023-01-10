package session

import (
	"openlou/ent"
	"time"

	"github.com/gorilla/websocket"
)

type Session struct {
	LastAction time.Time
	tries      int
	Conn       *websocket.Conn
	User       *ent.User
}

func (s *Session) NewTry() {
	s.tries++
}

func (s *Session) Close() {
	s.Conn.Close()
}
