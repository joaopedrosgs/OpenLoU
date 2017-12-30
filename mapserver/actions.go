package mapserver

import (
	"OpenLoU/entities"
	"encoding/json"
)

func triesToCreateCity(x int, y int, user_id int) bool {
	return true
}
func getCitiesJson(x int, y int) string {
	citiesJson, err := json.Marshal([10]entities.City{})
	if err != nil {
		return ""
	}
	return string(citiesJson)
}
