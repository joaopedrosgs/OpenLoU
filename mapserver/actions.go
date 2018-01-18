package mapserver

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/database"
	"github.com/joaopedrosgs/OpenLoU/entities"
	"strconv"
)

const (
	getCitiesQuery = "SELECT cities.continent_id, cities.name, cities.x, cities.y, cities.points, users.username FROM cities, users WHERE cities.continent_id = $1 AND cities.x >= $2 AND cities.x <= $3 AND cities.y >= $4 AND cities.y <= $5 AND cities.user_id = users.id LIMIT $6"
)

func (ms *mapserver) createCity(requestData map[string]string, answer *communication.Answer) {

}
func (ms *mapserver) getCities(requestData map[string]string, answer *communication.Answer) {

	x, err := strconv.Atoi(requestData["X"])
	if err != nil || x < 0 || x > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "Bad X value"
		return
	}
	y, err := strconv.Atoi(requestData["Y"])
	if err != nil || y < 0 || y > configuration.GetSingleton().Parameters.General.ContinentSize {
		answer.Data = "Bad Y value"
		return
	}
	distance, err := strconv.Atoi(requestData["Range"])
	if err != nil || distance <= 0 || distance > 10 {
		answer.Data = "Bad Range value"
		return
	}

	continent, err := strconv.Atoi(requestData["Continent"])
	if err != nil || continent < 0 || continent > 50 {
		answer.Data = "Bad Continent value"
		return
	}

	max_x := x + distance
	max_y := y + distance
	min_x := x - distance
	min_y := y - distance
	cities := []entities.City{}

	database.GetSingleton().Where("x BETWEEN ? AND ? AND y BETWEEN ? AND ?", min_x, max_x, min_y, max_y).Find(&cities)

	answer.Data = cities

}
