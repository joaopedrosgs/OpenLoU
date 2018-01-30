package worker

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
	log "github.com/sirupsen/logrus"
	"reflect"
	"runtime"
)

type Worker struct {
	endPoints  map[int]func(*communication.Request, *communication.Answer, *chan *communication.Answer)
	in         chan *communication.Request
	out        *chan *communication.Answer
	name       string
	LogContext *log.Entry
	majorCode  int
}

func (w *Worker) Setup(name string, majorCode int) {
	w.majorCode = majorCode
	w.name = name
	w.LogContext = log.WithFields(log.Fields{"Entity": w.name})
	w.in = make(chan *communication.Request)
	w.endPoints = make(map[int]func(*communication.Request, *communication.Answer, *chan *communication.Answer))

}
func (w *Worker) StartListening() {
	w.LogContext.Info(w.name + " started listening")

	for {
		request := <-w.in
		answer := request.ToAnswer()
		if endpoint, ok := w.endPoints[request.Type%100]; ok {
			go endpoint(request, answer, w.out)
		} else {
			*w.out <- answer
		}
	}
}
func (w *Worker) RegisterInternalEndpoint(endpoint func(*communication.Request, *communication.Answer, *chan *communication.Answer), minorCode int) {
	if _, exists := w.endPoints[minorCode]; !exists {
		log.WithFields(log.Fields{"Code": minorCode, "Name": runtime.FuncForPC(reflect.ValueOf(endpoint).Pointer()).Name()}).Info("New endpoint registered!")
		w.endPoints[minorCode] = endpoint
	} else {
		w.LogContext.WithField("Code", minorCode).Error("An endpoint with this minorCode already exists!")
	}
}

func (w *Worker) GetInChan() *chan *communication.Request {
	return &w.in
}
func (w *Worker) SetOutChan(out *chan *communication.Answer) {
	w.out = out
}
func (w *Worker) GetCode() int {
	return w.majorCode
}
func (w *Worker) GetName() string {
	return w.name
}
