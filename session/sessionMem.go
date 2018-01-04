package session

import (
	"OpenLoU/configuration"
	"time"
)

type sessionInfo struct {
	user_id     int
	last_action time.Time
	tries       int
}
type sessionMem struct {
	sessions map[string]*sessionInfo
}

func (s *sessionMem) NewSession(user_id int, key string) {
	s.sessions[key] = &sessionInfo{user_id, time.Now(), 0}
}

func (s *sessionMem) SessionExists(key string) bool {
	if len(key) != configuration.GetInstance().Parameters.Security.KeySize {
		return false
	}
	_, ok := s.sessions[key]
	return ok
}

func (s *sessionMem) DeleteSession(key string) {
	delete(s.sessions, key)
}

func NewSessionInMemory() *sessionMem {
	return &sessionMem{make(map[string]*sessionInfo)}
}

func (s *sessionMem) NewTry(key string) {
	s.sessions[key].tries++
}
