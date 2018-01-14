package communication

type Answer struct {
	userID   int
	key      string
	Ok       bool
	Type     int
	Data     interface{}
	isSystem bool
}

func (answer *Answer) GetId() int {
	return answer.userID
}
func (answer *Answer) GetKey() string {
	return answer.key
}
func (answer *Answer) SetUserID(id int) {
	answer.userID = id
}
func (answer *Answer) SerBySystem() {
	answer.isSystem = true
}

var badRequest = Answer{-1, "", false, -1, "Bad request", false}
var unauthorizedRequest = Answer{-1, "", false, -1, "Unauthorized request", false}

func BadRequest() *Answer {
	return &badRequest
}
func Unauthorized() *Answer {
	return &unauthorizedRequest
}

func (answer *Answer) IsSystem() bool {
	return answer.isSystem
}
