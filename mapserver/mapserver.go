package mapserver

import (
	"database/sql"
	"strconv"

	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	_ "github.com/lib/pq"
)

const (
	MAPSERVER            = 1
	GET_REGION_CITIES    = 100
	GET_USER_CITY_LIST   = 101
	GET_REGION_CITY_INFO = 102
	CREATE_CITY          = 103
)

type dataGetCities struct {
	X     int
	Y     int
	Range int
}

type dataCreateCity struct {
	X int
	Y int
}

type mapserver struct {
	database *sql.DB
	in       chan *communication.Request
	out      *chan *communication.Answer
	code     int
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
	if err != nil {
		return nil, err
	}

	return &mapserver{database, make(chan *communication.Request), nil, 1}, nil
}
func (ms *mapserver) StartListening() {
	println("Map server started listening")

	for {
		request := <-ms.in
		go ms.ProcessRequest(request)
	}
}
func (ms *mapserver) ProcessRequest(request *communication.Request) {
	answer := request.ToAnswer()
	answer.Data = make(map[string]string)
	switch request.Type {
	case GET_REGION_CITIES:
		{
			answer.Data["Result"] = "True"
			x, err := strconv.Atoi(request.Data["X"])
			if err != nil {
				break
			}
			y, err := strconv.Atoi(request.Data["Y"])
			if err != nil {
				break
			}
			answer.Data["Cities"] = getCitiesJson(x, y)

		}
	default:
	}
	*ms.out <- answer

}
