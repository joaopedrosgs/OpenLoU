package hub

import (
	"context"
	"errors"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/session"
	"github.com/joaopedrosgs/openlou/ent"
	"github.com/joaopedrosgs/openlou/ent/user"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var logger = log.WithFields(log.Fields{"Entity": "Hub"})
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Hub struct {
	servers []IServer
	workers map[int]*chan *communication.Request
	client  *ent.Client
}

func New() (*Hub, error) {
	hub := &Hub{}
	hub.workers = make(map[int]*chan *communication.Request)
	var err error
	hub.client, err = ent.Open("postgres", "dbname=postgres host=localhost user=postgres password=postgres sslmode=disable")
	if err != nil {
		logger.Error("Hub failed to connect to the database!")
		return nil, err
	}
	if err := hub.client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	logger.Info("Hub has been started!")
	return hub, nil

}

func (h *Hub) Start() {
	logger.Info("Hub is starting")

	for _, server := range h.servers {
		go server.StartListening()
	}

	http.HandleFunc("/api", h.handleUser)
	logger.Info("Hub started listening on " + viper.GetString("server.port"))

	err := http.ListenAndServe(viper.GetString("server.port"), nil)
	if err != nil {
		logger.Fatal(err)
	}

}
func (h *Hub) Authenticate(request *communication.Request, conn *websocket.Conn) (*session.Session, error) {

	username, exists := request.Data["username"]
	if !exists {
		return nil, errors.New("empty username")
	}

	password, exists := request.Data["password"]
	if !exists {
		return nil, errors.New("empty password")
	}

	user, err := h.client.User.Query().Where(user.NameEQ(username)).Only(context.Background())
	if err != nil {
		return nil, errors.New("wrong account info: " + err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("wrong account info: " + err.Error())
	}

	userSession, err := session.NewSession(user, conn)
	if err != nil {
		return nil, errors.New("failed to create sesion: " + err.Error())
	}

	return userSession, nil
}

func (h *Hub) handleUser(w http.ResponseWriter, r *http.Request) {
	log.Info("User connected")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn("Could not upgrade to Websocket because: ", err.Error())
		return
	}

	request := &communication.Request{}

	err = c.ReadJSON(request)
	if err != nil {
		log.Info("failed to read json: " + err.Error())
		c.Close()
		return
	}
	userSession, err := h.Authenticate(request, c)
	if err != nil {
		log.Info("failed to authenticate: ", err.Error())
		c.Close()
		return
	}
	request.SetSession(userSession)

	answer := request.ToAnswer()
	answer.Data = "Logged in!"
	answer.Result = true
	c.WriteJSON(answer)

	for {

		err := c.ReadJSON(request)
		if err != nil { // if the request could not be parsed
			log.Info("failed to read json:", err.Error())
			return
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

	server.SetClient(h.client)
	logger.WithFields(log.Fields{"Name": server.GetName(), "Endpoint": server.GetCode()}).Info("A server has been registered")
	h.servers = append(h.servers, server)
	server.AfterSetup()
	return nil
}
