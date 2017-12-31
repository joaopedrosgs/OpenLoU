package mapserver

import (
	"OpenLoU/communication"
	"OpenLoU/configuration"
	"database/sql"
	"fmt"
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

type Mapserver struct {
	database *sql.DB
	In       chan communication.Request
	Out      *chan communication.Answer
}

func CreateAndConnect(config *configuration.Config) (*Mapserver, error) {
	connectionString := "user=%s password=%s host=%s port=%d dbname=%s sslmode=%s"
	connectionString = fmt.Sprintf(connectionString, config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Name, config.Db.SSL)
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &Mapserver{database, make(chan communication.Request), nil}, nil
}
func (ms *Mapserver) StartListening() {
	println("Map server started listening")

	for {
		request := <-ms.In
		go ms.ProcessRequest(request)
	}
}
func (ms *Mapserver) ProcessRequest(request communication.Request) {
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
	*ms.Out <- answer

}

func (m *Mapserver) SetEndPoint(outChan *chan communication.Answer) {
	m.Out = outChan

}
