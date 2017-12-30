package mapserver

import (
	"OpenLoU/communication"
	"OpenLoU/configuration"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	GET_CITIES  = 100
	CREATE_CITY = 200
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
	Out      chan communication.Answer
}

func CreateAndConnect(config *configuration.Config) (*Mapserver, error) {
	connectionString := "user=%s password=%s host=%s port=%d dbname=%s sslmode=%s"
	connectionString = fmt.Sprintf(connectionString, config.Db.User, config.Db.Password, config.Db.Host, config.Db.Port, config.Db.Name, config.Db.SSL)
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &Mapserver{database, make(chan communication.Request), make(chan communication.Answer)}, nil
}
func (ms *Mapserver) StartListening() {
	for {
		request := <-ms.In
		go ms.ProcessRequest(request)
	}
}
func (ms *Mapserver) ProcessRequest(request communication.Request) {
	answer := communication.Answer{UserID: request.UserID, Type: request.Type, Result: false}
	switch request.Type {
	case GET_CITIES:
		{
			data := dataGetCities{}
			err := json.Unmarshal([]byte(request.Data), &data)
			if err == nil {
				answer.Result = true
				answer.Data = getCitiesJson(data.X, data.X)
			}

		}
	case CREATE_CITY:
		{
			data := dataCreateCity{}
			err := json.Unmarshal([]byte(request.Data), &data)
			if err == nil {
				answer.Result = triesToCreateCity(data.X, data.Y, request.UserID)
			}

		}
	default:
	}
	ms.Out <- answer

}
