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
)

const (
	MAPSERVER   = 2
	LOGINSERVER = 1
	CITYSERVER  = 3
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
	AuthorizedConnections   map[int]*net.Conn
	UnauthorizedConnections map[int]*net.Conn
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
	h.AuthorizedConnections = make(map[int]*net.Conn)
	h.UnauthorizedConnections = make(map[int]*net.Conn)
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
		if answer.Type%100 == LOGINSERVER && answer.Result == true {
			user, ok := h.UnauthorizedConnections[answer.UserID]
			if ok {
				delete(h.UnauthorizedConnections, answer.UserID)
				h.AuthorizedConnections[answer.UserID] = user
			}
		}
		userConn, ok := h.AuthorizedConnections[answer.UserID]
		if !ok {
			delete(h.AuthorizedConnections, answer.UserID)
			continue
		}
		write := bufio.NewWriter(*userConn)
		n := 0
		err := errors.New("")
		if answer.Result {
			n, err = write.Write([]byte(answer.Data))
		} else {
			n, err = write.Write([]byte(strconv.FormatBool(answer.Result)))
		}
		if err != nil || n == 0 {
			delete(h.AuthorizedConnections, answer.UserID)
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
			_, ok := h.AuthorizedConnections[request.UserID]
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
	h.UnauthorizedConnections[request.UserID] = conn
	*h.LoginChanIn <- request
}
func (h *Hermes) HandleAuthorizedUser(conn *net.Conn, request communication.Request) {
	server := request.Type % 100
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
