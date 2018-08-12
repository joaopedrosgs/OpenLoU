package hub

import (
	"errors"

	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"

	"net/http"

	"github.com/gorilla/websocket"
	"github.com/joaopedrosgs/OpenLoU/models"
	"github.com/joaopedrosgs/OpenLoU/session"
	"github.com/joaopedrosgs/OpenLoU/storage"
	log "github.com/sirupsen/logrus"
)

var context = log.WithFields(log.Fields{"Entity": "Hub"})
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Hub struct {
	servers  []IServer
	workers  map[int]*chan *models.Request
	connPool *pgx.ConnPool
	conn     *pgx.Conn
}

func New() (*Hub, error) {
	hub := &Hub{}
	hub.workers = make(map[int]*chan *models.Request)

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "127.0.0.1",
			Port:     5432,
			Database: "openlou",
			User:     "default",
		},
		MaxConnections: 5,
	}
	var err error
	hub.connPool, err = pgx.NewConnPool(poolConfig)
	if err != nil {
		context.Error("Hub failed to connect to the database!")
		return nil, err
	}
	hub.conn, err = hub.connPool.Acquire()
	if err != nil {
		context.Error("Hub failed to acquire connection!")
		return nil, err
	}
	context.Info("Hub has been started!")
	return hub, nil

}

func (h *Hub) Start(port string) {
	context.Info("Hub is starting")

	for _, server := range h.servers {
		go server.StartListening()
	}

	http.HandleFunc("/api", h.handleUser)
	context.Info("Hub started listening on " + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		context.Fatal(err)
	}

}
func (h *Hub) Authenticate(request *models.Request, conn *websocket.Conn) error {
	login, exists := request.Data["login"]
	if !exists {
		return errors.New("empty login")
	}
	password, exists := request.Data["password"]
	if !exists {
		return errors.New("empty password")
	}
	user, err := storage.GetUserInfo(h.conn, login)
	if err != nil {
		return errors.New("wrong account info: " + err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return errors.New("wrong account info: " + err.Error())
	}
	session, err := session.NewSession(user, conn)

	if err != nil {
		return errors.New("failed to create session: " + err.Error())
	}
	request.Session = session
	return nil
}

func (h *Hub) handleUser(w http.ResponseWriter, r *http.Request) {
	log.Info("User connected")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn("Could not upgrade to Websocket because: ", err.Error())
		return
	}
	request := &models.Request{}
	err = c.ReadJSON(request)

	if err != nil {
		log.Info("failed to read json: " + err.Error())
		c.Close()
		return
	}
	err = h.Authenticate(request, c)
	if err != nil {
		log.Info("failed to authenticate: ", err.Error())

		c.Close()
		return
	}
	answer := request.ToAnswer()
	answer.Data = "Logged in!"
	answer.Ok = true
	answer.SendToUser()
	for {
		err := c.ReadJSON(request)
		if err != nil { // if the request could not be parsed
			log.Info("failed to read json")
			continue

		}
		workerChan, ok := h.workers[request.EndPoint]
		if ok {
			*workerChan <- request
		}

	}
	c.Close()
}

func (h *Hub) RegisterServer(server IServer) error {
	if _, exists := h.workers[server.GetCode()]; exists {
		return errors.New("endpoint used by " + server.GetName())
	}
	h.workers[server.GetCode()] = server.GetJobsChan()
	conn, err := h.connPool.Acquire()
	if err != nil {
		context.Error("Failed to acquire connection from connection pool: ", err.Error)
		return err
	}
	server.SetConn(conn)
	context.WithFields(log.Fields{"Name": server.GetName(), "Endpoint": server.GetCode()}).Info("A server has been registered")
	h.servers = append(h.servers, server)
	return nil
}
