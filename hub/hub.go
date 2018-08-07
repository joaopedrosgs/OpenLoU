package hub

import (
	"errors"

	"github.com/jackc/pgx"

	"net"

	"flag"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/session"
	log "github.com/sirupsen/logrus"
)

var context = log.WithFields(log.Fields{"Entity": "Hub"})

const (
	MSGSIZE = 256
)

type Hub struct {
	listener       net.Listener
	servers        []IServer
	inChan         chan *communication.Answer
	workers        map[int]*chan *communication.Request
	connectionPool *pgx.ConnPool
}

func New() (*Hub, error) {
	hub := &Hub{}
	hub.workers = make(map[int]*chan *communication.Request)
	hub.inChan = make(chan *communication.Answer)
	connConfig := pgx.ConnConfig{Host: "127.0.0.1", Port: 5432, Database: "openlou", User: "default"}

	poolConfig := pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: 4}
	var err error
	hub.connectionPool, err = pgx.NewConnPool(poolConfig)
	if err != nil {
		context.Error("Hub failed to start!")
		return nil, err
	}
	context.Info("Hub has been started!")
	return hub, nil

}

var upgrader = websocket.Upgrader{}
var addr = flag.String("addr", "localhost:8080", "http service address")

func (h *Hub) Start(port string) {
	context.Info("Hub is starting")

	for _, server := range h.servers {
		go server.StartListening()
	}

	http.HandleFunc("/api", h.handleUser)
	context.Info("Hub started listening on " + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		context.Error(err)
	}

}

func (h *Hub) handleReturn() {
	var conn *websocket.Conn
	var ok bool
	for {
		answer := <-h.inChan
		conn, ok = session.GetUserConn(answer.GetKey())
		if ok {
			go h.writeBackToUser(answer, conn)
		}

	}
}
func (h *Hub) writeBackToUser(answer *communication.Answer, conn *websocket.Conn) {
	err := conn.WriteJSON(answer)
	if err != nil {
		if answer.IsSystem() {
			session.DeleteSessionByName(answer.GetUserName())
		} else {
			session.DeleteSession(answer.GetKey())

		}
		return
	}

}
func (h *Hub) handleUser(w http.ResponseWriter, r *http.Request) {
	log.Info("User connected")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn("Could not upgrade to Websocket")
	}
	defer c.Close()
	request := &communication.Request{}
	err = c.ReadJSON(request)

	if err != nil {
		h.writeBackToUser(communication.BadRequest(), c)
		return
	}
	err = h.SetupConn(request, c)

	if err != nil {
		log.Info("Invalid login:", err.Error())
		return
	}
	defer session.DeleteSession(request.Key)

	answer := &communication.Answer{}

	for {
		err := c.ReadJSON(request)
		if err != nil { // if the request could not be parsed
			answer = communication.BadRequest()
			h.writeBackToUser(answer, c)
			continue

		}
		if !session.Exists(request.Key) { // if the session doenst exist
			answer = communication.Unauthorized()
			h.writeBackToUser(answer, c)
			break
		}
		h.handleAuthorizedUser(request)

	}
}

func (h *Hub) SetupConn(request *communication.Request, conn *websocket.Conn) error {

	s, auth := session.GetSession(request.Key)
	if auth && s == nil {
		session.SetConn(request.Key, conn)
		return nil
	}

	return errors.New("session not found")
}
func (h *Hub) handleAuthorizedUser(request *communication.Request) {

	workerChan, ok := h.workers[request.EndPoint]
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

func (h *Hub) RegisterServer(server IServer) error {
	if _, exists := h.workers[server.GetCode()]; exists {
		return errors.New("endpoint used by " + server.GetName())
	}
	h.workers[server.GetCode()] = server.GetInChan()
	server.SetOutChan(&h.inChan)
	conn, err := h.connectionPool.Acquire()
	if err != nil {
		context.Error("Failed to acquire connection from connection pool: ", err.Error)
		return err
	}
	server.SetConn(conn)
	context.WithFields(log.Fields{"Name": server.GetName(), "Endpoint": server.GetCode()}).Info("A server has been registered")
	h.servers = append(h.servers, server)
	return nil
}
