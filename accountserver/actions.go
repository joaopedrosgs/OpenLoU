package accountserver

import (
	"github.com/joaopedrosgs/OpenLoU/models"
)

func (cs *accountServer) GetUserInfo(request *models.Request) *models.Answer {
	answer := request.ToAnswer()
	answer.Data = request.Session.User
	answer.Ok = true
	return answer

}
