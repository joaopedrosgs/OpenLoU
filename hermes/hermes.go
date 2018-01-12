// Hermes, the messenger
// Receives packages and passes them to the respective server (processor), receive the answer and pass them to user)

package hermes

import (
	"bufio"
	"encoding/json"
	"errors"

	"net"

	"github.com/joaopedrosgs/OpenLoU/communication"
	log "github.com/sirupsen/logrus"
	"strconv"
)

var context = log.WithFields(log.Fields{"Entity": "Hermes"})

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
		conn := h.sessions.GetUserConnByKey(answer.GetKey())
		go h.writeBackToUser(answer, conn)

	}
}
func (h *Hermes) writeBackToUser(answer *communication.Answer, conn net.Conn) {
	if conn == nil {
		context.Error("User connection is invalid")
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

		err := json.Unmarshal(buffer[:received], request)

		if err != nil { // failed do unmarshal
			h.writeBackToUser(communication.BadRequest(), conn)
		}

		if h.sessions.SessionExists(request.Key) {
			h.handleAuthorizedUser(request)
		} else {
			h.writeBackToUser(communication.Unauthorized(), conn)
		}
		received, err = reader.Read(buffer) // blocks until all the data is available

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
	servertype, err := strconv.Atoi(request.Data["Type"])
	if err != nil {
		server := servertype / 100
		workerChan, ok := h.workers[server]
		if ok {
			*workerChan <- request
		}
	} else {
		conn := h.sessions.GetUserConnByKey(request.Key)
		h.writeBackToUser(communication.BadRequest(), conn)
		h.sessions.NewTry(request.Key)
	}

}

func (h *Hermes) GetEntryPoint() *chan *communication.Answer {
	return &h.inChan
}

func (h *Hermes) RegisterWorker(worker IWorker) error {
	if _, exists := h.workers[worker.GetCode()]; exists {
		return errors.New("Code already used")
	}
	h.workers[worker.GetCode()] = worker.GetInChan()
	worker.SetOutChan(&h.inChan)
	context.WithField("Code", worker.GetCode()).Info("A worker has been registered")
	return nil
}
