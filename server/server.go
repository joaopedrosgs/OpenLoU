package server

import (
	"openlou/ent"
	"reflect"
	"runtime"

	log "github.com/sirupsen/logrus"
	"openlou/communication"
)

type Server struct {
	endPoints               map[int]func(*communication.Request) *communication.Answer
	jobs                    chan *communication.Request
	name                    string
	LogContext              *log.Entry
	EndPointCode            int
	internalWorkerInstances int
	client                  *ent.Client
}

func (w *Server) Setup(name string, endPointCode int, workerInstances int) {
	w.EndPointCode = endPointCode
	w.name = name
	w.LogContext = log.WithFields(log.Fields{"Entity": w.name})
	w.jobs = make(chan *communication.Request)
	w.endPoints = make(map[int]func(*communication.Request) *communication.Answer)
	w.internalWorkerInstances = workerInstances

}
func (w *Server) SetClient(client *ent.Client) {
	w.client = client
}
func (w *Server) GetClient() *ent.Client {
	return w.client
}
func (w *Server) StartListening() {
	w.LogContext.Info(w.name + " started listening")
	for i := 0; i < w.internalWorkerInstances; i++ {
		go w.Worker()
	}
}
func (w *Server) Worker() {
	for request := range w.jobs {
		if endpoint, ok := w.endPoints[request.Type]; ok {
			answer := endpoint(request)
			answer.Dispatch()
		} else {
			request.ToAnswer().Dispatch()
		}
	}
}
func (w *Server) RegisterInternalEndpoint(endpoint func(*communication.Request) *communication.Answer, minorCode int) {
	if _, exists := w.endPoints[minorCode]; !exists {
		log.WithFields(log.Fields{"Code": minorCode, "Name": runtime.FuncForPC(reflect.ValueOf(endpoint).Pointer()).Name()}).Info("New endpoint registered!")
		w.endPoints[minorCode] = endpoint
	} else {
		w.LogContext.WithField("Code", minorCode).Error("An endpoint with this type already exists!")
	}
}

func (w *Server) GetJobsChan() *chan *communication.Request {
	return &w.jobs
}
func (w *Server) GetCode() int {
	return w.EndPointCode
}
func (w *Server) GetName() string {
	return w.name
}
