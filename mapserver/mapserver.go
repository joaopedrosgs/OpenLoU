package mapserver

import (
	"database/sql"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"time"
)

var context = log.WithFields(log.Fields{"Entity": "Map Server"})

type mapserver struct {
	database  *sql.DB
	endPoints map[int]func(*communication.Request, *communication.Answer)
	in        chan *communication.Request
	out       *chan *communication.Answer
	code      int
}

func (ms *mapserver) GetInChan() *chan *communication.Request {
	return &ms.in
}

func (ms *mapserver) SetOutChan(out *chan *communication.Answer) {
	ms.out = out
}

func (ms *mapserver) GetCode() int {
	return ms.code
}

func New() (*mapserver, error) {

	database, err := sql.Open("postgres", configuration.GetConnectionString())
	for err != nil {
		context.Error("Failed to connect to db: " + err.Error())
		context.Info("Trying again in 10 seconds...")
		time.Sleep(10 * time.Second)
		database, err = sql.Open("postgres", configuration.GetConnectionString())

	}

	ms := &mapserver{database, make(map[int]func(*communication.Request, *communication.Answer)), make(chan *communication.Request), nil, 1}
	ms.registerInternalEndpoint(triesToCreateCity, 101)
	ms.registerInternalEndpoint(getCities, 102)
	return ms, nil

}
func (ms *mapserver) StartListening() {
	context.Info("Map server started listening")

	for {
		request := <-ms.in
		go ms.ProcessRequest(request)
	}
}
func (ms *mapserver) ProcessRequest(request *communication.Request) {
	answer := request.ToAnswer()
	endpoint, ok := ms.endPoints[request.Type]
	if ok {
		endpoint(request, answer)
	}
	*ms.out <- answer

}

func (ms *mapserver) registerInternalEndpoint(endpoint func(*communication.Request, *communication.Answer), code int) {
	if _, exists := ms.endPoints[code]; !exists {
		ms.endPoints[code] = endpoint
	}
}
