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

type SessionMem struct {
	sessions map[string]*Session
}

func (s *SessionMem) NewSession(user_id int, key string) bool {
	if len(key) == configuration.GetInstance().Parameters.Security.KeySize || user_id >= 0 {
		if !s.SessionExistsByID(user_id) {
			s.sessions[key] = &Session{user_id, time.Now(), 0, nil}
			return true
		}
	}
	return false
}

func (s *SessionMem) SetConn(key string, conn net.Conn) {
	session, ok := s.sessions[key]
	if ok {
		session.conn = conn
	}
}

func (s *SessionMem) SessionExists(key string) bool {
	_, ok := s.sessions[key]
	return ok
}
func (s *SessionMem) SessionExistsByID(id int) bool {
	for _, session := range s.sessions {
		if session.user_id == id {
			return true
		}
	}
	return false
}

func (s *SessionMem) DeleteSession(key string) {
	delete(s.sessions, key)

}

func (s *SessionMem) DeleteSessionByID(id int) {
	for key, session := range s.sessions {
		if session.user_id == id {
			delete(s.sessions, key)
		}
	}

}
func NewSessionInMemory() *SessionMem {
	return &SessionMem{make(map[string]*Session)}
}

func (s *SessionMem) NewTry(key string) {
	if session, ok := s.sessions[key]; ok {
		session.tries++
	}
}

func (s *SessionMem) GetSession(key string) (*Session, bool) {
	session, ok := s.sessions[key]
	return session, ok
}
func (s *SessionMem) GetSessionById(id int) (*Session, bool) {
	for _, session := range s.sessions {
		if session.user_id == id {
			return session, true
		}
	}
	return nil, false
}

func (s *SessionMem) GetUserConn(key string) (net.Conn, bool) {
	session, ok := s.sessions[key]
	if ok {
		return session.conn, ok
	}
	return nil, ok

}
func (s *SessionMem) GetUserConnById(id int) (net.Conn, bool) {
	for _, session := range s.sessions {
		if session.user_id == id {
			return session.conn, true
		}
	}
	return nil, false

}
