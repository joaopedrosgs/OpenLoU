package session

import (
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"net"
	"sync"
	"time"
)

type Session struct {
	user_id     uint
	last_action time.Time
	tries       int
	conn        net.Conn
}

type SessionMem struct {
	mutex    sync.RWMutex
	sessions map[string]*Session
}

func (s *SessionMem) NewSession(user_id uint, key string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(key) == configuration.GetSingleton().Parameters.Security.KeySize || user_id >= 0 {
		if !s.sessionsExistsByID(user_id) {
			s.sessions[key] = &Session{user_id, time.Now(), 0, nil}
			return true
		}
	}
	return false
}

func (s *SessionMem) SetConn(key string, conn net.Conn) {
	s.mutex.Lock()
	session, ok := s.sessions[key]
	if ok {
		session.conn = conn
	}
	s.mutex.Unlock()
}

func (s *SessionMem) SessionExists(key string) bool {
	s.mutex.RLock()
	_, ok := s.sessions[key]
	s.mutex.RUnlock()
	return ok
}
func (s *SessionMem) sessionsExistsByID(id uint) bool {
	for _, session := range s.sessions {
		if session.user_id == id {
			return true
		}
	}
	return false
}

func (s *SessionMem) DeleteSession(key string) {
	s.mutex.Lock()
	delete(s.sessions, key)
	s.mutex.Unlock()

}

func (s *SessionMem) DeleteSessionByID(id uint) {
	s.mutex.Lock()
	for key, session := range s.sessions {
		if session.user_id == id {
			delete(s.sessions, key)
		}
	}
	s.mutex.Unlock()

}
func NewSessionInMemory() *SessionMem {
	return &SessionMem{sync.RWMutex{}, make(map[string]*Session)}
}

func (s *SessionMem) NewTry(key string) {
	s.mutex.RLock()
	if session, ok := s.sessions[key]; ok {
		session.tries++
	}
	s.mutex.RUnlock()
}

func (s *SessionMem) GetUserConn(key string) (net.Conn, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	session, ok := s.sessions[key]
	if ok {
		return session.conn, ok
	}
	return nil, ok

}
func (s *SessionMem) GetUserConnById(id uint) (net.Conn, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for _, session := range s.sessions {
		if session.user_id == id {
			return session.conn, true
		}
	}
	return nil, false

}
