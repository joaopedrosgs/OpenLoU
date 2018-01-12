package mapserver

import (
	"encoding/json"
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/entities"
	"strconv"
)

const (
	getCitiesQuery = "SELECT cities.continent_id, cities.name, cities.x, cities.y, cities.points, users.username FROM cities, users WHERE cities.continent_id = $1 AND cities.x >= $2 AND cities.x <= $3 AND cities.y >= $4 AND cities.y <= $5 AND cities.user_id = users.id"
)

func (ms *mapserver) triesToCreateCity(request *communication.Request, answer *communication.Answer) {

}
func (ms *mapserver) getCities(request *communication.Request, answer *communication.Answer) {
	x, err := strconv.Atoi(request.Data["X"])
	if err != nil {
		answer.Data = "Bad X value"
	}
	y, err := strconv.Atoi(request.Data["Y"])
	if err != nil {
		answer.Data = "Bad Y value"
		return
	}
	distance, err := strconv.Atoi(request.Data["Range"])
	if err != nil {
		answer.Data = "Bad Range value"
		return
	}
	continent, err := strconv.Atoi(request.Data["Continent"])
	if err != nil {
		answer.Data = "Bad Continent value"
		return
	}
	query := ms.database.QueryRow(getCitiesQuery, continent, x-distance, x+distance, y-distance, y+distance)
}
