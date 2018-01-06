package communication

type Answer struct {
	internal_id int32
	key         string
	Type        int
	Data        map[string]string
	isSystem    bool
}

func (answer *Answer) GetKey() string {
	return answer.key
}

var badRequest = Answer{0, "", -1, map[string]string{"Result": "False", "Message": "Bad request"}, false}
var unauthorizedRequest = Answer{0, "", -1, map[string]string{"Result": "False", "Message": "Bad request"}, false}

func BadRequest() *Answer {
	return &badRequest
}
func Unauthorized() *Answer {
	return &unauthorizedRequest
}

func (answer *Answer) IsSystem() bool {
	return answer.isSystem
}

func (answer *Answer) GetInternalID() int32 {
	return answer.internal_id
}
