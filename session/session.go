package session

import (
	"OpenLoU/configuration"
	"database/sql"
	"fmt"
	"time"
)

type sessionDB struct {
	db *sql.DB
}

const (
	//table name
	tableSessions = "sessions"

	//Queries
	newSessionQuery    = "INSERT INTO " + tableSessions + "(user_id, key, last_action, tries) VALUES ($1, $2, $3, $4)"
	sessionExistsQuery = "SELECT EXISTS (SELECT 1 FROM " + tableSessions + " WHERE key=$1)"
	deleteSessionQuery = "DELETE from " + tableSessions + " WHERE key=$1"
)

func (s *sessionDB) NewSession(user_id int, key string) {
	if user_id >= 0 {
		_, err := s.db.Query(newSessionQuery, user_id, key, time.Now(), 0)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}

func (s *sessionDB) SessionExists(key string) bool {
	result := false
	if key != "" {
		err := s.db.QueryRow(sessionExistsQuery, key).Scan(&result)
		if err != nil {
			println(err.Error())
		}
	}
	return result
}

func (s *sessionDB) DeleteSession(key string) {
	s.db.Query(deleteSessionQuery, key)
}

func New() *sessionDB {
	database, err := sql.Open("postgres", configuration.GetConnectionString())
	if err != nil {
		return nil
	}
	return &sessionDB{database}
}
