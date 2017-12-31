// Hermes, the messenger
// Receives packages and passes them to the respective server (processor), receive the answer and pass them to user)

package hermes

import (
	"OpenLoU/communication"
	"bufio"
	"encoding/json"
	"errors"
	"net"
	"sync"
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

	// Login Server requests always start with 3
	LOGINSERVER = 3
)

type IAnswer interface {
	ToJson() string
	GetUserId() int
}
type IRequest interface {
	GetType() int
	GetUserId() int
}

type Hermes struct {
	listener          net.Listener
	Connections       map[string]*communication.Connection
	UnauthConnections map[*net.Conn]*communication.Connection
	MapChanIn         *chan communication.Request
	LoginChanIn       *chan communication.Request
	CityChanIn        *chan communication.Request
	InChan            chan communication.Answer
	Mutex             sync.RWMutex
}

func Create(mapChanIn *chan communication.Request, loginChanIn *chan communication.Request, cityChanIn *chan communication.Request) Hermes {
	h := Hermes{}
	h.MapChanIn = mapChanIn
	h.LoginChanIn = loginChanIn
	h.CityChanIn = cityChanIn
	h.InChan = make(chan communication.Answer)
	h.Connections = make(map[string]*communication.Connection)
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
	go h.HandleReturn()
	for {
		client, err := h.listener.Accept()
		if err == nil {
			go h.HandleUser(&client)
		}
	}
}

func (h *Hermes) HandleReturn() {
	for {
		answer := <-h.InChan
		userConn := answer.GetConn()
		if answer.Type/100 == LOGINSERVER && answer.Result == true { // if the login went Ok
			h.Mutex.Lock()
			userConn.Key = answer.Data
			h.Connections[userConn.Key] = answer.GetConn() // Como recuperar a conexáo se ela nao foi inserida?
			h.Mutex.Unlock()

		}

		if userConn != nil {
			h.writeBackToUser(userConn, answer)
		}
	}
}
func (h *Hermes) writeBackToUser(userConn *communication.Connection, answer communication.Answer) {
	if userConn.Conn == nil {
		println("User connection is invalid")
		return
	}
	n := 0
	err := errors.New("")
	if answer.Result {
		n, err = userConn.Writer.Write([]byte(answer.Data))
	} else {
		res, _ := json.Marshal(answer)
		n, err = userConn.Writer.Write([]byte(res))
	}
	if err != nil || n == 0 {
		h.Mutex.Lock()
		delete(h.Connections, answer.GetKey())
		h.Mutex.Unlock()
		return
	}
	userConn.Writer.Flush()
}
func (h *Hermes) HandleUser(conn *net.Conn) {
	println("Hermes handled a new user")
	tries := 0
	user := &communication.Connection{conn, bufio.NewWriter(*conn), make([]byte, 1024), bufio.NewReader(*conn), make([]byte, 1024), ""}
	request := &communication.Request{}
	request.SetConn(user)

	for {
		received, err := user.Reader.Read(user.BufferRead)
		if err != nil || tries > 5 {
			break
		} else {
			err := json.Unmarshal(user.BufferRead[:received], request)
			if err != nil {
				h.writeBackToUser(user, communication.BadRequest())
				tries++
				continue
			}
			println(request.Data["Login"])
			h.Mutex.RLock()
			_, ok := h.Connections[request.Key]
			h.Mutex.RUnlock()
			if ok && request.Key != "" {
				println("Calling Auth request")
				h.HandleAuthorizedUser(*request)
			} else {
				println("Calling unauth request")
				h.HandleUnauthorizedUser(user, request)
			}

		}
	}

	(*conn).Close()
}
func (h *Hermes) HandleUnauthorizedUser(conn *communication.Connection, request *communication.Request) {
	*h.LoginChanIn <- *request
}
func (h *Hermes) HandleAuthorizedUser(request communication.Request) {
	server := request.Type / 100
	switch server {
	case MAPSERVER:
		{
			println("Going to mapchan")
			*h.MapChanIn <- request
		}
	case CITYSERVER:
		{
			println("Going to citychan")
			*h.CityChanIn <- request
		}
	}

}
