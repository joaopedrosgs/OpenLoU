package communication

type Answer struct {
	userID   uint
	key      string
	Ok       bool
	Type     int
	Data     interface{} `json:",omitempty"`
	isSystem bool
}

func (answer *Answer) GetId() uint {
	return answer.userID
}
func (answer *Answer) GetKey() string {
	return answer.key
}
func (answer *Answer) SetUserID(id uint) {
	answer.userID = id
}
func (answer *Answer) SetBySystem() {
	answer.isSystem = true
}

var badRequest = Answer{0, "", false, -1, "Bad request", false}
var unauthorizedRequest = Answer{0, "", false, -1, "Unauthorized request", false}

func BadRequest() *Answer {
	return &badRequest
}
func Unauthorized() *Answer {
	return &unauthorizedRequest
}

func (answer *Answer) IsSystem() bool {
	return answer.isSystem
}
