package worker

import (
	"github.com/joaopedrosgs/OpenLoU/hub"
	log "github.com/sirupsen/logrus"
	"reflect"
	"runtime"
)

type Worker struct {
	endPoints  map[int]func(request *hub.RequestWithCallback)
	in         chan *hub.RequestWithCallback
	name       string
	LogContext *log.Entry
	majorCode  int
}

func (w *Worker) Setup(name string, majorCode int) {
	w.majorCode = majorCode
	w.name = name
	w.LogContext = log.WithFields(log.Fields{"Entity": w.name})
	w.in = make(chan *hub.RequestWithCallback)
	w.endPoints = make(map[int]func(request *hub.RequestWithCallback))

}
func (w *Worker) StartListening() {
	w.LogContext.Info(w.name + " started listening")

	for {
		received := <-w.in
		if endpoint, ok := w.endPoints[received.Request.Type%100]; ok {
			go endpoint(received)
		} else {
			go received.Callback(received.Request.ToAnswer())
		}
	}
}
func (w *Worker) RegisterInternalEndpoint(endpoint func(request *hub.RequestWithCallback), minorCode int) {
	if _, exists := w.endPoints[minorCode]; !exists {
		log.WithFields(log.Fields{"Code": minorCode, "Name": runtime.FuncForPC(reflect.ValueOf(endpoint).Pointer()).Name()}).Info("New endpoint registered!")
		w.endPoints[minorCode] = endpoint
	} else {
		w.LogContext.WithField("Code", minorCode).Error("An endpoint with this minorCode already exists!")
	}
}

func (w *Worker) GetInChan() *chan *hub.RequestWithCallback {
	return &w.in
}
func (w *Worker) GetCode() int {
	return w.majorCode
}
func (w *Worker) GetName() string {
	return w.name
}
