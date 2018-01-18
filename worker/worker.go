package worker

import (
	"database/sql"
	"github.com/joaopedrosgs/OpenLoU/communication"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"reflect"
	"runtime"
)

type Worker struct {
	database   *sql.DB
	endPoints  map[int]func(map[string]string, *communication.Answer)
	in         chan *communication.Request
	out        *chan *communication.Answer
	name       string
	LogContext *log.Entry
	code       int
}

func (w *Worker) Setup(name string, code int) {
	w.code = code
	w.name = name
	w.LogContext = log.WithFields(log.Fields{"Entity": w.name})
	w.in = make(chan *communication.Request)
	w.endPoints = make(map[int]func(map[string]string, *communication.Answer))

}
func (w *Worker) StartListening() {
	w.LogContext.Info(w.name + " started listening")

	for {
		request := <-w.in
		go w.ProcessRequest(request)
	}
}
func (w *Worker) ProcessRequest(request *communication.Request) {
	answer := request.ToAnswer()
	endpoint, ok := w.endPoints[request.Type]
	if ok {
		answer.Ok = true
		endpoint(request.Data, answer)

	}
	*w.out <- answer

}
func (w *Worker) RegisterInternalEndpoint(endpoint func(map[string]string, *communication.Answer), code int) {
	if _, exists := w.endPoints[code]; !exists {
		log.WithFields(log.Fields{"Code": code, "Name": runtime.FuncForPC(reflect.ValueOf(endpoint).Pointer()).Name()}).Info("New endpoint registered")
		w.endPoints[code] = endpoint
	}
}

func (w *Worker) GetInChan() *chan *communication.Request {
	return &w.in
}
func (w *Worker) SetOutChan(out *chan *communication.Answer) {
	w.out = out
}
func (w *Worker) GetCode() int {
	return w.code
}

func (w *Worker) GetDatabase() *sql.DB {
	return w.database
}

func (w *Worker) GetName() string {
	return w.name
}
