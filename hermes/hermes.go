// Hermes, the messenger
// Receives packages and passes them to the respective server (processor), receive the answer and pass them to user)

package hermes

import (
	"OpenLoU/communication"
	"bufio"
	"encoding/json"
	"errors"
	"net"
	"strconv"
	"sync"
)

const (
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
	listener                net.Listener
	AuthorizedConnections   sync.Map
	UnauthorizedConnections sync.Map
	MapChanIn               *chan communication.Request
	LoginChanIn             *chan communication.Request
	CityChanIn              *chan communication.Request
	InChan                  chan communication.Answer
}

func Create(mapChanIn *chan communication.Request, loginChanIn *chan communication.Request, cityChanIn *chan communication.Request) Hermes {
	h := Hermes{}
	h.MapChanIn = mapChanIn
	h.LoginChanIn = loginChanIn
	h.CityChanIn = cityChanIn
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
		if answer.Type/100 == LOGINSERVER && answer.Result == true {
			user, ok := h.UnauthorizedConnections.Load(answer.UserID)
			if ok {
				h.UnauthorizedConnections.Delete(answer.UserID)
				h.AuthorizedConnections.Store(answer.UserID, user)
			}

		}

		userConn, ok := h.AuthorizedConnections.Load(answer.UserID)
		if !ok {

			h.AuthorizedConnections.Delete(answer.UserID)
			continue
		}
		write := bufio.NewWriter(userConn.(net.Conn))
		n := 0
		err := errors.New("")
		if answer.Result {
			n, err = write.Write([]byte(answer.Data))
		} else {
			n, err = write.Write([]byte(strconv.FormatBool(answer.Result)))
		}
		if err != nil || n == 0 {
			h.AuthorizedConnections.Delete(answer.UserID)
			continue
		}
	}

}
func (h *Hermes) HandleUser(conn *net.Conn) {
	println("Hermes handled a new user")

	tries := 0

	buffer := make([]byte, 1024)
	read := bufio.NewReader(*conn)
	request := communication.Request{UserID: -1}

	for {
		received, err := read.Read(buffer)
		if err != nil || tries > 5 {
			break
		} else {
			err := json.Unmarshal(buffer[:received], request)
			if err != nil {
				tries++
				break
			}
			_, ok := h.AuthorizedConnections.Load(request.UserID)
			if !ok {
				h.HandleUnauthorizedUser(conn, request)
			} else {
				h.HandleAuthorizedUser(conn, request)
			}
		}
	}

	(*conn).Close()
}
func (h *Hermes) HandleUnauthorizedUser(conn *net.Conn, request communication.Request) {
	h.UnauthorizedConnections.Store(request.UserID, conn)
	*h.LoginChanIn <- request
}
func (h *Hermes) HandleAuthorizedUser(conn *net.Conn, request communication.Request) {
	server := request.Type / 100
	switch server {
	case MAPSERVER:
		{
			*h.MapChanIn <- request
		}
	case CITYSERVER:
		{
			*h.CityChanIn <- request
		}
	}

}
