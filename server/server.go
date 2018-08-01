package server

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
	log "github.com/sirupsen/logrus"
	"reflect"
	"runtime"
)

type Server struct {
	endPoints               map[int]func(*communication.Request) *communication.Answer
	in                      chan *communication.Request
	out                     *chan *communication.Answer
	name                    string
	LogContext              *log.Entry
	majorCode               int
	internalWorkerInstances int
}

func (w *Server) Setup(name string, majorCode int, workerInstances int) {
	w.majorCode = majorCode
	w.name = name
	w.LogContext = log.WithFields(log.Fields{"Entity": w.name})
	w.in = make(chan *communication.Request)
	w.endPoints = make(map[int]func(*communication.Request) *communication.Answer)
	w.internalWorkerInstances = workerInstances

}
func (w *Server) StartListening() {
	w.LogContext.Info(w.name + " started listening")
	for i := 0; i < w.internalWorkerInstances; i++ {
		go w.Worker()
	}
}
func (w *Server) Worker() {
	for request := range w.in {
		answer := request.ToAnswer()
		if endpoint, ok := w.endPoints[request.Type%100]; ok {
			*w.out <- endpoint(request)
		} else {
			*w.out <- answer
		}
	}
}
func (w *Server) RegisterInternalEndpoint(endpoint func(*communication.Request) *communication.Answer, minorCode int) {
	if _, exists := w.endPoints[minorCode]; !exists {
		log.WithFields(log.Fields{"Code": minorCode, "Name": runtime.FuncForPC(reflect.ValueOf(endpoint).Pointer()).Name()}).Info("New endpoint registered!")
		w.endPoints[minorCode] = endpoint
	} else {
		w.LogContext.WithField("Code", minorCode).Error("An endpoint with this minorCode already exists!")
	}
}

func (w *Server) GetInChan() *chan *communication.Request {
	return &w.in
}
func (w *Server) SetOutChan(out *chan *communication.Answer) {
	w.out = out
}
func (w *Server) GetCode() int {
	return w.majorCode
}
func (w *Server) GetName() string {
	return w.name
}
