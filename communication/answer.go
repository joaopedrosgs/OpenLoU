package communication

type Answer struct {
	userName string
	key      string
	Ok       bool
	Type     int
	Data     interface{} `json:",omitempty"`
	isSystem bool
}

func (answer *Answer) GetUserName() string {
	return answer.userName
}
func (answer *Answer) GetKey() string {
	return answer.key
}
func (answer *Answer) SetUserName(userName string) {
	answer.userName = userName
}
func (answer *Answer) SetBySystem() {
	answer.isSystem = true
}

var badRequest = Answer{"", "", false, -1, "Bad request", false}
var unauthorizedRequest = Answer{"", "", false, -1, "Unauthorized request", false}

func BadRequest() *Answer {
	return &badRequest
}
func Unauthorized() *Answer {
	return &unauthorizedRequest
}

func (answer *Answer) IsSystem() bool {
	return answer.isSystem
}
