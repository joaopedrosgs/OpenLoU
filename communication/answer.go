package communication

type Answer struct {
	key      string
	Ok       bool
	Type     int
	Data     string
	isSystem bool
}

func (answer *Answer) GetKey() string {
	return answer.key
}

var badRequest = Answer{"", false, -1, "Bad request", false}
var unauthorizedRequest = Answer{"", false, -1, "Unauthorized request", false}

func BadRequest() *Answer {
	return &badRequest
}
func Unauthorized() *Answer {
	return &unauthorizedRequest
}

func (answer *Answer) IsSystem() bool {
	return answer.isSystem
}
