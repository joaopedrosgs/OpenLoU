package modules

type Module interface {
	GetType() string
	GetId() int
}
