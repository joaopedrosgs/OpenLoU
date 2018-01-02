package mapserver

import (
	"OpenLoU/communication"
	"OpenLoU/configuration"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
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
}

func New() (*mapserver, error) {

	database, err := sql.Open("postgres", configuration.GetConnectionString())
	if err != nil {
		return nil, err
	}

	return &mapserver{database, make(chan *communication.Request), nil}, nil
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
	switch request.Type {
	case GET_REGION_CITIES:
		{

			answer.Result = true
			x, err := strconv.Atoi(request.Data["X"])
			if err != nil {
				break
			}
			y, err := strconv.Atoi(request.Data["Y"])
			if err != nil {
				break
			}
			answer.Data = getCitiesJson(x, y)

		}
	default:
	}
	*ms.out <- answer

}

func (m *mapserver) SetEndPoint(outChan *chan *communication.Answer) {
	m.out = outChan

}

func (m *mapserver) GetEntryPoint() *chan *communication.Request {
	return &m.in

}
