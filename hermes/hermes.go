// Hermes, the messenger
// Receives packages and passes them to the respective server (processor), receive the answer and pass them to user)

package hermes

import (
	"bufio"
	"encoding/json"
	"errors"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"net"
)

const (
	MSG_SIZE = 1024
	HERMES   = 0
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

type Hermes struct {
	listener net.Listener
	inChan   chan *communication.Answer
	sessions ISessionBackend
	workers  map[int]*chan *communication.Request
}

func Create(backend ISessionBackend) Hermes {
	h := Hermes{}
	h.workers = make(map[int]*chan *communication.Request)
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
			go h.handleUser(client)
		}
	}
}

func (h *Hermes) handleReturn() {
	for {
		answer := <-h.inChan
		conn := h.sessions.GetUserConnByKey(answer.GetKey())
		go h.writeBackToUser(answer, conn)

	}
}
func (h *Hermes) writeBackToUser(answer *communication.Answer, conn net.Conn) {
	if conn == nil {
		println("User connection is invalid")
		h.sessions.DeleteSession(answer.GetKey())
		return
	}
	n := 0
	err := errors.New("")

	buffer, _ := json.Marshal(answer)
	writer := bufio.NewWriter(conn)
	n, err = writer.Write(buffer)

	if err != nil || n == 0 {
		h.sessions.DeleteSession(answer.GetKey())
		return
	}
	writer.Flush()
}
func (h *Hermes) handleUser(conn net.Conn) {
	defer conn.Close()
	request := &communication.Request{}
	reader := bufio.NewReader(conn)
	buffer := make([]byte, MSG_SIZE)
	received, err := reader.Read(buffer) // blocks until all the data is available

	if err == nil {
		json.Unmarshal(buffer[:received], request)
	}
	if h.AuthRequest(request, conn) {

		for {
			if err == nil && received > 0 && received < MSG_SIZE {

				err := json.Unmarshal(buffer[:received], request)

				if err != nil { // failed do unmarshal
					h.writeBackToUser(communication.BadRequest(), conn)
				}

				if h.sessions.SessionExists(request.Key) {
					h.handleAuthorizedUser(request)
				} else {
					h.writeBackToUser(communication.Unauthorized(), conn)
				}

			} else {
				break
			}
			received, err = reader.Read(buffer) // blocks until all the data is available
		}
	}
	h.sessions.DeleteSession(request.Key)

}

func (h *Hermes) AuthRequest(request *communication.Request, conn net.Conn) bool {
	auth := false

	auth = h.sessions.SessionExists(request.Key)
	if auth {
		h.sessions.SetConn(request.Key, conn)
	}

	return auth
}
func (h *Hermes) handleAuthorizedUser(request *communication.Request) {
	server := request.Type / 100
	workerChan, ok := h.workers[server]
	if ok {
		*workerChan <- request
	} else {
		conn := h.sessions.GetUserConnByKey(request.Key)
		h.writeBackToUser(communication.BadRequest(), conn)
		h.sessions.NewTry(request.Key)
	}

}

func (h *Hermes) GetEntryPoint() *chan *communication.Answer {
	return &h.inChan
}

func (h *Hermes) RegisterWorker(worker IWorker) {
	h.workers[worker.GetCode()] = worker.GetInChan()
	worker.SetOutChan(&h.inChan)
	println("Entity: Hermes, Message: A worker has been registered")
}
