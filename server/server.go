package server

import (
	"reflect"
	"runtime"

	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	endPoints               map[int]func(*models.Request) *models.Answer
	jobs                    chan *models.Request
	name                    string
	LogContext              *log.Entry
	EndPointCode            int
	internalWorkerInstances int
	conn                    *pgx.Conn
}

func (w *Server) Setup(name string, endPointCode int, workerInstances int) {
	w.EndPointCode = endPointCode
	w.name = name
	w.LogContext = log.WithFields(log.Fields{"Entity": w.name})
	w.jobs = make(chan *models.Request)
	w.endPoints = make(map[int]func(*models.Request) *models.Answer)
	w.internalWorkerInstances = workerInstances

}
func (w *Server) SetConn(conn *pgx.Conn) {
	w.conn = conn
}
func (w *Server) GetConn() *pgx.Conn {
	return w.conn
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
			answer.SendToUser()
		} else {
			request.ToAnswer().SendToUser()
		}
	}
}
func (w *Server) RegisterInternalEndpoint(endpoint func(*models.Request) *models.Answer, minorCode int) {
	if _, exists := w.endPoints[minorCode]; !exists {
		log.WithFields(log.Fields{"Code": minorCode, "Name": runtime.FuncForPC(reflect.ValueOf(endpoint).Pointer()).Name()}).Info("New endpoint registered!")
		w.endPoints[minorCode] = endpoint
	} else {
		w.LogContext.WithField("Code", minorCode).Error("An endpoint with this type already exists!")
	}
}

func (w *Server) GetJobsChan() *chan *models.Request {
	return &w.jobs
}
func (w *Server) GetCode() int {
	return w.EndPointCode
}
func (w *Server) GetName() string {
	return w.name
}
