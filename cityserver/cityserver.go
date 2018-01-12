package cityserver

import (
	"database/sql"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"time"
)

var context = log.WithFields(log.Fields{"Entity": "City Server"})

type cityserver struct {
	database *sql.DB
	in       chan *communication.Request
	out      *chan *communication.Answer
	code     int
}

func (cs *cityserver) GetInChan() *chan *communication.Request {
	return &cs.in
}

func (cs *cityserver) SetOutChan(out *chan *communication.Answer) {
	cs.out = out
}

func (cs *cityserver) GetCode() int {
	return cs.code
}

func New() (*cityserver, error) {

	database, err := sql.Open("postgres", configuration.GetConnectionString())
	for err != nil {
		context.Error("Failed to connect to db: " + err.Error())
		context.Info("Trying again in 10 seconds...")
		time.Sleep(10 * time.Second)
		database, err = sql.Open("postgres", configuration.GetConnectionString())

	}

	return &cityserver{database, make(chan *communication.Request), nil, 2}, nil
}
func (cs *cityserver) StartListening() {
	context.Info("City server started listening")

	for {
		request := <-cs.in
		go cs.ProcessRequest(request)
	}
}
func (cs *cityserver) ProcessRequest(request *communication.Request) {
	answer := request.ToAnswer()
	*cs.out <- answer

}
