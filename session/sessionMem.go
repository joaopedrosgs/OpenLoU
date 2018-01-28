package session

import (
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"net"
	"sync"
	"time"
)

type Session struct {
	userId     uint
	lastAction time.Time
	tries      int
	conn       net.Conn
}

type sessionMem struct {
	mutex    sync.RWMutex
	sessions map[string]*Session
}

var sessionsStorage sessionMem

func NewSession(userId uint, key string) bool {
	sessionsStorage.mutex.Lock()
	if len(key) == configuration.GetSingleton().Parameters.Security.KeySize || userId >= 0 {
		sessionsStorage.sessions[key] = &Session{userId, time.Now(), 0, nil}
		sessionsStorage.mutex.Unlock()
		return true

	}
	sessionsStorage.mutex.Unlock()
	return false
}

func SetConn(key string, conn net.Conn) {
	sessionsStorage.mutex.Lock()
	session, ok := sessionsStorage.sessions[key]
	if ok {
		session.conn = conn
	}
	sessionsStorage.mutex.Unlock()
}

func Exists(key string) bool {
	sessionsStorage.mutex.RLock()
	_, ok := sessionsStorage.sessions[key]
	sessionsStorage.mutex.RUnlock()
	return ok
}
func sessionsExistsByID(id uint) bool {
	for _, session := range sessionsStorage.sessions {
		if session.userId == id {
			return true
		}
	}
	return false
}

func DeleteSession(key string) {
	sessionsStorage.mutex.Lock()
	delete(sessionsStorage.sessions, key)
	sessionsStorage.mutex.Unlock()

}

func DeleteSessionByID(id uint) {
	sessionsStorage.mutex.Lock()
	for key, session := range sessionsStorage.sessions {
		if session.userId == id {
			delete(sessionsStorage.sessions, key)
		}
	}
	sessionsStorage.mutex.Unlock()

}
func NewSessionInMemory() {
	sessionsStorage = sessionMem{sync.RWMutex{}, make(map[string]*Session)}
}

func NewTry(key string) {
	sessionsStorage.mutex.RLock()
	if session, ok := sessionsStorage.sessions[key]; ok {
		session.tries++
	}
	sessionsStorage.mutex.RUnlock()
}

func GetUserConn(key string) (net.Conn, bool) {
	sessionsStorage.mutex.RLock()
	session, ok := sessionsStorage.sessions[key]
	if ok {
		sessionsStorage.mutex.RUnlock()
		return session.conn, ok

	}
	sessionsStorage.mutex.RUnlock()
	return nil, ok

}
func GetUserConnById(id uint) (net.Conn, bool) {
	sessionsStorage.mutex.RLock()
	for _, session := range sessionsStorage.sessions {
		if session.userId == id {
			sessionsStorage.mutex.RUnlock()
			return session.conn, true
		}
	}
	sessionsStorage.mutex.RUnlock()
	return nil, false

}

func GetUserID(key string) uint {
	sessionsStorage.mutex.RLock()
	id := sessionsStorage.sessions[key].userId
	sessionsStorage.mutex.RUnlock()

	return id
}
