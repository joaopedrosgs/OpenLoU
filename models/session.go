package models

import (
	"time"

	"github.com/gorilla/websocket"
)

type Session struct {
	LastAction time.Time
	tries      int
	Conn       *websocket.Conn
	User       *User
}

func (s *Session) NewTry() {
	s.tries++
}

func (s *Session) Close() {
	s.Conn.Close()
}
