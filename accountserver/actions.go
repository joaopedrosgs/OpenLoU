package accountserver

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/storage"
)

func (cs *accountServer) GetUserInfo(request *communication.Request) *communication.Answer {
	answer := request.ToAnswer()
	user, err := storage.GetUserInfoByKey(cs.GetConn(), request.Key)
	if err != nil {
		cs.LogContext.WithField("When", "Retrieving user information").Error(err.Error())
		return nil
	}
	answer.Data = user
	answer.Ok = true
	return answer

}
