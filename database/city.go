package database

import (
	"errors"
	"github.com/joaopedrosgs/OpenLoU/entities"
)

func GetUserCities(userID uint) ([]entities.City, error) {
	rows, err := db.Query("SELECT id, x, y, continent_id, type, name, points FROM cities WHERE user_id = $1", userID)
	defer rows.Close()
	if err != nil {
		return nil, errors.New("Could not pull cities: " + err.Error())
	}
	var cities []entities.City
	for rows.Next() {
		city := entities.City{}
		err := rows.Scan(&city.ID, &city.X, &city.Y, &city.ContinentID, &city.Type, &city.Name, &city.Points)
		if err != nil {
			return nil, errors.New("Failed to scan: " + err.Error())
		}
		cities = append(cities, city)
	}
	if err != nil {
		return nil, errors.New("Could not pull cities: " + err.Error())
	}

	return cities, nil
}

func GetCitiesInRange(x, y, radius, continent uint) (*[]entities.City, error) {
	cities := make([]entities.City, 0, (radius*radius)*4)
	if x < radius {
		x = radius
	}
	if y < radius {
		y = radius
	}
	minX := x - radius
	minY := y - radius
	maxX := x + radius
	maxY := y + radius
	rows, err := db.Query("SELECT id, x, y, continent_id, type, name, points "+
		"FROM cities WHERE "+
		"x BETWEEN $1 AND $2 AND "+
		"y BETWEEN $3 AND $4 "+
		"AND continent_id = $5;", minX, maxX, minY, maxY, continent)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		city := entities.City{}
		err := rows.Scan(&city.ID, &city.X, &city.Y, &city.ContinentID, &city.Type, &city.Name, &city.Points)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	return &cities, nil
}

func createCity(userID, x, y, continentID uint) error {
	var id uint
	err := db.QueryRow("INSERT INTO cities(user_id, x, y, continent_id) VALUES($1, $2, $3, $4) RETURNING id", userID, x, y, continentID).Scan(&id)
	if err != nil {
		return errors.New("Failed to create city: " + err.Error())
	}
	err = setTile(continentID, x, y, "city")
	if err != nil {
		return errors.New("Failed to create city: " + err.Error())
	}
	err = CreateConstruction(id, 11, 11, 1, 1)
	if err != nil {
		return errors.New("Failed to create city: " + err.Error())
	}
	_, err = db.Exec("UPDATE continents SET  number_of_cities = number_of_cities + 1 WHERE id = $1", continentID)
	if err != nil {
		return errors.New("Failed to create city: " + err.Error())
	}

	return nil
}

func findNewCityLocation() (uint, uint, uint, error) {
	var randX, randY, continentID uint

	err := db.QueryRow("SELECT id from continents WHERE is_active limit 1").Scan(&continentID)
	if err != nil {
		context.WithField("When", "Finding a continent").Error(err.Error())
		return 0, 0, 0, err
	}
	err = db.QueryRow("SELECT x,y "+
		"FROM tiles "+
		"WHERE occupied_by = 'land' AND "+
		"continent_id = $1 "+
		"ORDER by random() limit 1", continentID).Scan(&randX, &randY)
	if err != nil {
		context.WithField("When", "Finding tile").Error(err.Error())
		return 0, 0, 0, err
	}
	return randX, randY, continentID, nil
}
