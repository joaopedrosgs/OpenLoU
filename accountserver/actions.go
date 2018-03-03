package accountserver

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/database"
)

func (cs *accountServer) GetUserInfo(request *communication.Request, answer *communication.Answer) *communication.Answer {
	if user, err := database.GetUserInfoByKey(request.Key); err != nil {
		cs.LogContext.WithField("When", "Retrieving user information").Error(err.Error())
	} else {
		answer.Data = user
		answer.Ok = true
	}
	return answer

}
