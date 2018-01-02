// Hermes, the messenger
// Receives packages and passes them to the respective server (processor), receive the answer and pass them to user)

package hermes

import (
	"OpenLoU/communication"
	"bufio"
	"encoding/json"
	"errors"
	"net"
)

const (
	HERMES = 0
	// Map Server requests always start with 1
	MAPSERVER            = 1
	GET_REGION_CITIES    = 100
	GET_USER_CITY_LIST   = 101
	GET_REGION_CITY_INFO = 102

	// City Server requests always start with 2
	CITYSERVER             = 2
	GET_CITY_CONSTRUCTIONS = 200
	GET_CITY_INFO          = 201
	UPGRADE_CONSTRUCTION   = 202
	DEGRADE_CONSTRUCTION   = 203

	// Email Server requests always start with 3
	LOGINSERVER = 3
)

type ISessionBackend interface {
	NewSession(id int, key string)

	SessionExists(key string) bool

	DeleteSession(key string)
}

type Hermes struct {
	listener   net.Listener
	mapChanIn  *chan *communication.Request
	cityChanIn *chan *communication.Request
	inChan     chan *communication.Answer
	sessions   ISessionBackend
}

func Create(mapChanIn *chan *communication.Request, backend ISessionBackend) Hermes {
	h := Hermes{}
	h.mapChanIn = mapChanIn
	h.inChan = make(chan *communication.Answer)
	h.sessions = backend
	println("Hermes has been started!")
	return h

}

func (h *Hermes) StartListening() {
	err := errors.New("")
	h.listener, err = net.Listen("tcp", ":8080")
	if err != nil {
		println("Hermes has failed to listen: " + err.Error())
		return
	}
	println("Hermes started listening")
	go h.handleReturn()
	for {
		client, err := h.listener.Accept()
		if err == nil {
			go h.handleUser(&client)
		}
	}
}

func (h *Hermes) handleReturn() {
	for {
		answer := <-h.inChan
		exists := h.sessions.SessionExists(answer.GetKey())
		h.writeBackToUser(answer.GetConn(), answer)
		if !exists {
			answer.GetConn().Close()
		}
	}
}
func (h *Hermes) writeBackToUser(userConn *communication.User, answer *communication.Answer) {
	if userConn.Conn == nil {
		println("User connection is invalid")
		return
	}
	n := 0
	err := errors.New("")
	res, _ := json.Marshal(answer)
	n, err = userConn.Writer.Write([]byte(res))

	if (err != nil || n == 0) && answer.GetKey() != "" {
		h.sessions.DeleteSession(answer.GetKey())
		return
	}
	userConn.Writer.Flush()
}
func (h *Hermes) handleUser(conn *net.Conn) {
	println("Hermes handled a new user")
	tries := 0
	user := &communication.User{conn, bufio.NewWriter(*conn), make([]byte, 1024), bufio.NewReader(*conn), make([]byte, 1024), ""}
	request := &communication.Request{}
	request.SetConn(user)

	for {
		received, err := user.Reader.Read(user.BufferRead)
		if err != nil || tries > 5 {
			break // conexao caiu ou mtas tentativas
		} else {
			err := json.Unmarshal(user.BufferRead[:received], request)
			if err != nil {
				h.writeBackToUser(user, communication.BadRequest())
				tries++
				continue
			}
			ok := h.sessions.SessionExists(request.Key)
			if ok {
				user.Key = request.Key
				h.handleAuthorizedUser(request)
			} else {
				h.writeBackToUser(user, communication.Unauthorized())
			}

		}
	}
	ok := h.sessions.SessionExists(user.Key)
	if ok {
		h.sessions.DeleteSession(user.Key)
	}
	(*conn).Close()
}

func (h *Hermes) handleAuthorizedUser(request *communication.Request) {
	server := request.Type / 100
	switch server {
	case MAPSERVER:
		{
			println("Going to mapchan")
			*h.mapChanIn <- request
		}
	case CITYSERVER:
		{
			println("Going to citychan")
			*h.cityChanIn <- request
		}
	default:
		h.writeBackToUser(request.GetConn(), communication.BadRequest())
	}

}

func (h *Hermes) GetEntryPoint() *chan *communication.Answer {
	return &h.inChan
}
