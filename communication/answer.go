package communication

import (
	"openlou/session"
)

type Answer struct {
	Result  bool
	Type    int
	Data    interface{} `json:",omitempty"`
	session *session.Session
}

func (a *Answer) Dispatch() {
	a.session.Conn.WriteJSON(a)
}
