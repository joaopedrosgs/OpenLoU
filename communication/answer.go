package communication

type Answer struct {
	key      string
	Type     int
	Result   bool
	Data     string
	isSystem bool
	connptr  *Connection
}

func (answer *Answer) GetKey() string {
	return answer.key
}

func BadRequest() Answer {
	return Answer{"", -1, false, "Bad request", false, nil}
}

func (answer *Answer) IsSystem() bool {
	return answer.isSystem
}

func (answer *Answer) GetConn() *Connection {
	return answer.connptr
}
