package hub

type IWorker interface {
	GetName() string
	GetInChan() *chan *RequestWithCallback
	GetCode() int
}
