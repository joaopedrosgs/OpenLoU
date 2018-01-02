package communication

type Answer struct {
	key      string
	Type     int
	Result   bool
	Data     string
	isSystem bool
	connptr  *User
}

func (answer *Answer) GetKey() string {
	return answer.key
}

var badRequest = Answer{"", -1, false, "Bad request", false, nil}
var unauthorizedRequest = Answer{"", -1, false, "Unauthorized", false, nil}

func BadRequest() *Answer {
	return &badRequest
}
func Unauthorized() *Answer {
	return &unauthorizedRequest
}

func (answer *Answer) IsSystem() bool {
	return answer.isSystem
}

func (answer *Answer) GetConn() *User {
	return answer.connptr
}
