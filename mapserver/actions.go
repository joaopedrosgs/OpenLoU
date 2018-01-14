package mapserver

import (
	"github.com/joaopedrosgs/OpenLoU/communication"
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
	if err != nil {
		answer.Data = "Bad X value"
		return
	}
	y, err := strconv.Atoi(requestData["Y"])
	if err != nil {
		answer.Data = "Bad Y value"
		return
	}
	distance, err := strconv.Atoi(requestData["Range"])
	if err != nil {
		answer.Data = "Bad Range value"
		return
	}
	continent, err := strconv.Atoi(requestData["Continent"])
	if err != nil {
		answer.Data = "Bad Continent value"
		return
	}
	query, err := ms.GetDatabase().Query(getCitiesQuery, continent, x-distance, x+distance, y-distance, y+distance, distance)
	if err != nil {
		answer.Data = "Failed to query"
		return
	}
	distance = distance + 2
	data := make([]*entities.City, 0, distance*distance)
	defer query.Close()
	i := 0
	for query.Next() {
		city := &entities.City{}
		query.Scan(&city.ContinentID, &city.Name, &city.X, &city.Y, &city.Points, &city.UserName)
		data = append(data, city)
		i++
	}
	answer.Data = data

}
