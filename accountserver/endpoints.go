package accountserver

import (
	core "github.com/joaopedrosgs/loucore/pkg"
	"github.com/joaopedrosgs/openlou/communication"
)

func (cs *accountServer) GetUserInfo(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()

	answer.Data = ""
	answer.Result = true
	return answer

}
func (cs *accountServer) CreateAccount(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()

	err := request.FieldsExist("email", "name", "password")
	if err != nil {
		return answer
	}
	user, err := core.CreateAccount(request.Data["email"], request.Data["name"], request.Data["password"])
	if err != nil {
		return answer
	}
	answer.Data = "Success"
	answer.Result = true
	request.GetSession().User = user
	return answer

}
func (cs *accountServer) Login(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()

	err := request.FieldsExist("email", "password")
	if err != nil {
		return answer
	}
	user, err := core.Authentication(request.Data["email"], request.Data["password"])
	if err != nil {
		return answer
	}
	answer.Data = "Success"
	answer.Result = true
	request.GetSession().User = user
	return answer

}
