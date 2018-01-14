// Hermes, the messenger
// Receives packages and passes them to the respective server (processor), receive the answer and pass them to user)

package hermes

import (
	"bufio"
	"encoding/json"
	"errors"

	"net"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/session"
	log "github.com/sirupsen/logrus"
)

var context = log.WithFields(log.Fields{"Entity": "Hermes"})

const (
	MSG_SIZE = 1024
)

type Hermes struct {
	listener net.Listener
	inChan   chan *communication.Answer
	sessions *session.SessionMem
	workers  map[int]*chan *communication.Request
}

func Create(backend *session.SessionMem) Hermes {
	h := Hermes{}
	h.workers = make(map[int]*chan *communication.Request)
	h.inChan = make(chan *communication.Answer)
	h.sessions = backend
	context.Info("Hermes has been started!")
	return h

}

func (h *Hermes) StartListening() {
	err := errors.New("")
	h.listener, err = net.Listen("tcp", ":8080")
	if err != nil {
		context.Error("Hermes has failed to listen: " + err.Error())
		return
	}
	context.Info("Hermes started listening")
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
		var conn net.Conn
		var ok bool
		if answer.IsSystem() {
			conn, ok = h.sessions.GetUserConnById(answer.GetId())
		} else {
			conn, ok = h.sessions.GetUserConn(answer.GetKey())
		}
		if ok {
			go h.writeBackToUser(answer, conn)
		}

	}
}
func (h *Hermes) writeBackToUser(answer *communication.Answer, conn net.Conn) {

	buffer, _ := json.Marshal(answer)
	writer := bufio.NewWriter(conn)
	n, err := writer.Write(buffer)

	if err != nil || n == 0 {
		if answer.IsSystem() {
			h.sessions.DeleteSessionByID(answer.GetId())
		} else {
			h.sessions.DeleteSession(answer.GetKey())

		}
		return
	}
	writer.Flush()

}
func (h *Hermes) handleUser(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)
	buffer := make([]byte, MSG_SIZE)
	received, err := reader.Read(buffer) // blocks until all the data is available

	if err != nil {
		h.writeBackToUser(communication.BadRequest(), conn)
		return
	}

	request := &communication.Request{}
	err = json.Unmarshal(buffer[:received], request)

	defer h.sessions.DeleteSession(request.Key)

	if err != nil {
		h.writeBackToUser(communication.BadRequest(), conn)
		return
	}

	if !h.Validate(request, conn) {
		h.writeBackToUser(communication.Unauthorized(), conn)
	}

	for err == nil && received > 0 && received < MSG_SIZE {

		if h.sessions.SessionExists(request.Key) {
			h.handleAuthorizedUser(request)
		} else {
			h.writeBackToUser(communication.Unauthorized(), conn)
			break
		}
		received, err = reader.Read(buffer) // blocks until all the data is available
		err := json.Unmarshal(buffer[:received], request)

		if err != nil { // failed do unmarshal
			answer := request.ToAnswer()
			answer.Data = err.Error()
			h.writeBackToUser(answer, conn)
			break
		}

	}
}

func (h *Hermes) Validate(request *communication.Request, conn net.Conn) bool {
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
		conn, ok := h.sessions.GetUserConn(request.Key)
		if ok {
			h.writeBackToUser(communication.BadRequest(), conn)
			h.sessions.NewTry(request.Key)
		}
	}

}

func (h *Hermes) GetEntryPoint() *chan *communication.Answer {
	return &h.inChan
}

func (h *Hermes) RegisterWorker(worker IWorker) error {
	if _, exists := h.workers[worker.GetCode()]; exists {
		return errors.New("Code used by " + worker.GetName())
	}
	h.workers[worker.GetCode()] = worker.GetInChan()
	worker.SetOutChan(&h.inChan)
	context.WithFields(log.Fields{"Name": worker.GetName(), "Code": worker.GetCode()}).Info("A worker has been registered")
	return nil
}
