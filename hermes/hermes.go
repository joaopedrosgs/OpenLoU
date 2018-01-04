// Hermes, the messenger
// Receives packages and passes them to the respective server (processor), receive the answer and pass them to user)

package hermes

import (
	"OpenLoU/communication"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sync/atomic"
	"time"
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

var request_internal_id int32 = -1

type Hermes struct {
	listener net.Listener
	inChan   chan *communication.Answer
	sessions ISessionBackend
	stats    map[int32]time.Time
	workers  map[int]*chan *communication.Request
}

func Create(backend ISessionBackend) Hermes {
	h := Hermes{}
	h.workers = make(map[int]*chan *communication.Request)
	h.inChan = make(chan *communication.Answer)
	h.sessions = backend
	h.stats = make(map[int32]time.Time)
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
		exists := h.sessions.SessionExists(answer.GetKey())
		h.writeBackToUser(answer)
		if !exists {
			answer.GetConn()
		}

	}
}
func (h *Hermes) writeBackToUser(answer *communication.Answer) {
	if answer.GetConn() == nil {
		println("User connection is invalid")
		return
	}
	n := 0
	err := errors.New("")

	buffer, _ := json.Marshal(answer)
	n, err = answer.GetWriter().Write(buffer)

	if (err != nil || n == 0) && answer.GetKey() != "" {
		h.sessions.DeleteSession(answer.GetKey())
		return
	}
	answer.GetWriter().Flush()
	fmt.Println(answer.GetInternalID(), " - ", (time.Now().Sub(h.stats[answer.GetInternalID()])).Round(time.Millisecond), "Result: ", answer.Result)
}
func (h *Hermes) handleUser(conn net.Conn) {
	defer conn.Close()
	request := &communication.Request{}
	atomic.AddInt32(&request_internal_id, 1)
	request.SetInternalID(request_internal_id)
	request.SetConn(conn)
	h.stats[request.GetInternalID()] = time.Now()

	for {
		received, err := request.Reader.Read(request.BufferRead) // blocks until all the data is available

		if err != nil { // failed to read
			break
		} else {

			err := json.Unmarshal(request.BufferRead[:received], request)

			if err != nil { // faiiled do unmarshal
				h.writeBackToUser(communication.BadRequest(conn))
				break
			}
			ok := h.sessions.SessionExists(request.Key)

			if ok {
				h.handleAuthorizedUser(request)
			} else {
				h.writeBackToUser(communication.Unauthorized(conn))
			}

		}

	}
	exists := h.sessions.SessionExists(request.Key)
	if exists {
		h.sessions.DeleteSession(request.Key)
	}

}

func (h *Hermes) handleAuthorizedUser(request *communication.Request) {
	server := request.Type / 100
	workerChan, ok := h.workers[server]
	if ok {
		*workerChan <- request
	} else {
		h.writeBackToUser(communication.BadRequest(request.GetConn()))
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
