// Hermes, the messenger
// Receives packages and passes them to the respective server (processor), receive the answer and pass them to user)

package hub

import (
	"bufio"
	"encoding/json"
	"errors"

	"net"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/session"
	log "github.com/sirupsen/logrus"
)

var context = log.WithFields(log.Fields{"Entity": "Hub"})

const (
	MSGSIZE = 256
)

type Hub struct {
	listener net.Listener
	inChan   chan *communication.Answer
	workers  map[int]*chan *communication.Request
}

func Create() Hub {
	h := Hub{}
	h.workers = make(map[int]*chan *communication.Request)
	h.inChan = make(chan *communication.Answer)
	context.Info("Hub has been started!")
	return h

}

func (h *Hub) StartListening(port string) {
	err := errors.New("")
	h.listener, err = net.Listen("tcp", port)
	if err != nil {
		context.Error("Hub has failed to listen: " + err.Error())
		return
	}
	context.Info("Hub started listening")
	go h.handleReturn()
	for {
		client, err := h.listener.Accept()
		if err == nil {
			go h.handleUser(client)
		}
	}
}

func (h *Hub) handleReturn() {
	var conn net.Conn
	var ok bool
	for {
		answer := <-h.inChan
		if answer.IsSystem() {
			conn, ok = session.GetUserConnById(answer.GetId())
		} else {
			conn, ok = session.GetUserConn(answer.GetKey())
		}
		if ok {
			go h.writeBackToUser(answer, conn)
		}

	}
}
func (h *Hub) writeBackToUser(answer *communication.Answer, conn net.Conn) {
	buffer, _ := json.Marshal(answer)
	writer := bufio.NewWriterSize(conn, MSGSIZE)
	n, err := writer.Write(buffer)

	if err != nil || n == 0 {
		if answer.IsSystem() {
			session.DeleteSessionByID(answer.GetId())
		} else {
			session.DeleteSession(answer.GetKey())

		}
		return
	}
	writer.Flush()

}
func (h *Hub) handleUser(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReaderSize(conn, MSGSIZE)
	buffer := make([]byte, MSGSIZE)
	received, err := reader.Read(buffer) // blocks until all the data is available

	if err != nil || received > MSGSIZE {
		h.writeBackToUser(communication.BadRequest(), conn)
		return
	}

	request := &communication.Request{}
	err = json.Unmarshal(buffer[:received], request)

	if err != nil || !h.Validate(request, conn) {
		h.writeBackToUser(communication.BadRequest(), conn)
		return
	}
	defer session.DeleteSession(request.Key)

	for err == nil && received > 0 && received < MSGSIZE {

		if session.Exists(request.Key) {
			h.handleAuthorizedUser(request)
		} else {
			h.writeBackToUser(communication.Unauthorized(), conn)
			break
		}

		received, err = reader.Read(buffer) // blocks until all the data is available
		if err != nil {
			answer := request.ToAnswer()
			answer.Data = err.Error()
			h.writeBackToUser(answer, conn)
			break
		}
		err := json.Unmarshal(buffer[:received], request)

		if err != nil { // failed do unmarshal
			answer := request.ToAnswer()
			answer.Data = err.Error()
			h.writeBackToUser(answer, conn)
			break
		}

	}
}

func (h *Hub) Validate(request *communication.Request, conn net.Conn) bool {
	auth := false

	auth = session.Exists(request.Key)
	if auth {
		session.SetConn(request.Key, conn)
	}

	return auth
}
func (h *Hub) handleAuthorizedUser(request *communication.Request) {

	server := request.Type / 100
	workerChan, ok := h.workers[server]
	if ok {
		*workerChan <- request
	} else {
		conn, ok := session.GetUserConn(request.Key)
		if ok {
			h.writeBackToUser(communication.BadRequest(), conn)
			session.NewTry(request.Key)
		}
	}

}

func (h *Hub) GetEntryPoint() *chan *communication.Answer {
	return &h.inChan
}

func (h *Hub) RegisterWorker(worker IWorker) error {
	if _, exists := h.workers[worker.GetCode()]; exists {
		return errors.New("Code used by " + worker.GetName())
	}
	h.workers[worker.GetCode()] = worker.GetInChan()
	worker.SetOutChan(&h.inChan)
	context.WithFields(log.Fields{"Name": worker.GetName(), "Code": worker.GetCode()}).Info("A worker has been registered")
	return nil
}
