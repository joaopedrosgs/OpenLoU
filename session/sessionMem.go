package session

import (
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"net"
	"time"
)

type Session struct {
	user_id     int
	last_action time.Time
	tries       int
	conn        net.Conn
}

type sessionMem struct {
	sessions map[string]*Session
}

func (s *sessionMem) NewSession(user_id int, key string) {
	if len(key) == configuration.GetInstance().Parameters.Security.KeySize || user_id >= 0 {
		s.sessions[key] = &Session{user_id, time.Now(), 0, nil}
	}
}

func (s *sessionMem) SetConn(key string, conn net.Conn) {
	if s.sessions[key].conn == nil {
		s.sessions[key].conn = conn
	}
}

func (s *sessionMem) SessionExists(key string) bool {
	if len(key) != configuration.GetInstance().Parameters.Security.KeySize {
		return false
	}
	_, ok := s.sessions[key]
	return ok
}

func (s *sessionMem) DeleteSession(key string) {
	if s.SessionExists(key) {
		delete(s.sessions, key)
	}
}

func NewSessionInMemory() *sessionMem {
	return &sessionMem{make(map[string]*Session)}
}

func (s *sessionMem) NewTry(key string) {
	s.sessions[key].tries++
}

func (s *sessionMem) GetUserId(key string) int {
	return s.sessions[key].user_id
}

func (s *sessionMem) GetUserConnByKey(key string) net.Conn {
	return s.sessions[key].conn
}

func (s *sessionMem) GetSession(key string) *Session {
	return s.sessions[key]
}

func (s *sessionMem) GetUserConnByID(id int) net.Conn {
	for _, v := range s.sessions {
		if v.user_id == id {
			return v.conn
		}
	}
	return nil

}
