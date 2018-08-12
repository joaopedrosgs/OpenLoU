package models

type Answer struct {
	session *Session
	Ok      bool
	Type    int
	Data    interface{} `json:",omitempty"`
}

func (answer *Answer) GetSession() *Session {
	return answer.session
}
func (answer *Answer) SendToUser() error {
	return answer.GetSession().Conn.WriteJSON(answer)
}
