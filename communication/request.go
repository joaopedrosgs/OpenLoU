package communication

import (
	"fmt"
	"github.com/joaopedrosgs/openlou/session"
	"time"
)

type Request struct {
	receivedAt time.Time
	EndPoint   int
	Type       int
	Data       map[string]string
	session    *session.Session
}

func (r *Request) FieldsExist(fields ...string) error {
	for _, field := range fields {
		if _, ok := r.Data[field]; !ok {
			return fmt.Errorf("Empty %s value!", field)
		}

	}
	return nil
}
func (r *Request) ToAnswer() *Answer {
	return &Answer{false, r.Type, "Bad Request", r.session}
}
func (r *Request) SetSession(userSession *session.Session) {
	if r.session == nil {
		r.session = userSession
	}
}
func (r *Request) GetSession() *session.Session {
	return r.session
}
