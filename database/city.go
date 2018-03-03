package database

import (
	"github.com/joaopedrosgs/OpenLoU/entities"
	"math/rand"
	"time"
)

func GetUserCities(userName string) (*[]entities.City, error) {
	var cities []entities.City
	if rows, err := db.Query(
		"SELECT x, y, continent_x, continent_y, type, name, points "+
			"FROM cities "+
			"WHERE user_name = $1", userName); err != nil {
		context.Error(err.Error())
	} else {
		cities = make([]entities.City, 0, 100)
		for rows.Next() {
			city := entities.City{UserName: userName}
			err := rows.Scan(&city.X, &city.Y, &city.ContinentX, &city.ContinentY, &city.Type, &city.Name, &city.Points)
			if err != nil {
				context.Error(err.Error())
				break
			}
			cities = append(cities, city)

		}
		rows.Close()

	}
	return &cities, nil
}

func GetCitiesInRange(x, y, radius uint) (*[]entities.City, error) {
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
	rows, err := db.Query(
		"SELECT x, y, continent_x, continent_y, type, name, points "+
			"FROM cities "+
			"WHERE x BETWEEN $1 AND $2 AND y BETWEEN $3 AND $4",
		minX, maxX, minY, maxY)
	if err != nil {
		context.WithField("When", "Finding cities in the region").Error(err.Error())
	} else {
		for rows.Next() {
			city := entities.City{}
			err := rows.Scan(&city.X, &city.Y, &city.ContinentX, &city.ContinentY, &city.Type, &city.Name, &city.Points)
			if err != nil {
				context.WithField("When", "Scaning cities from the database").Error(err.Error())
				break
			}
			cities = append(cities, city)
		}
	}
	rows.Close()
	return &cities, err
}

func createCity(city entities.City) error {
	_, err := db.Exec(
		"INSERT INTO cities(user_name, x, y, continent_x, continent_y) "+
			"VALUES($1, $2, $3, $4, $5)",
		city.UserName, city.X, city.Y, city.ContinentX, city.ContinentY)
	if err != nil {
		context.WithField("When", "Inserting city on the database").Error(err.Error())
	} else if err = setTile(city.TileNode); err != nil {
		context.WithField("When", "Setting the tile of the city").Error(err.Error())
	} else if err = PopulateNewCity(city); err != nil {
		context.WithField("When", "Trying to populate the city").Error(err.Error())
	} else if _, err = db.Exec(
		"UPDATE continents "+
			"SET number_of_cities = number_of_cities + 1 "+
			"WHERE x = $1 and y = $2",
		city.ContinentX, city.ContinentY); err != nil {
		context.WithField("When", "Setting the number of cities in the continent").Error(err.Error())
	}

	return err
}
func PopulateNewCity(city entities.City) error {
	var err error
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	var defaultConstruction = make([]entities.Construction, 0, 21)
	defaultConstruction[0] = entities.Construction{
		X:          11,
		Y:          11,
		CityX:      city.X,
		CityY:      city.Y,
		Type:       5,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Level:      1,
		Production: 300,
		Modifier:   1,
	}
	for i := 1; i < 20; i++ {
		var resourceType = random.Uint32() % 4
		var resourceX = random.Uint32() % 21
		var resourceY = random.Uint32() % 21
		defaultConstruction[i] = entities.Construction{
			X:          uint(resourceX),
			Y:          uint(resourceY),
			CityX:      city.X,
			CityY:      city.Y,
			Type:       uint(resourceType),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Level:      1,
			Production: 0,
			Modifier:   0,
		}
	}
	for i := 0; i < len(defaultConstruction); i++ {
		err = CreateConstruction(&defaultConstruction[i])
		if err != nil {
			context.WithField("When", "Creating default constructions").Error(err.Error())
		}
	}
	return err
}

func findNewCityLocation() (uint, uint, uint, uint, error) {
	var x, y, continentX, continentY uint
	var err error
	if err := db.QueryRow(
		"SELECT x,y from continents"+
			" WHERE is_active "+
			"limit 1").Scan(&continentX, &continentY); err != nil {
		context.WithField("When", "Finding a continent").Error(err.Error())
	} else if err = db.QueryRow(
		"SELECT x,y "+
			"FROM tiles "+
			"WHERE occupied_by = 'land' AND "+
			"continent_x = $1 AND "+
			"continent_y = $2 "+
			"ORDER by random() "+
			"limit 1", continentX, continentY).Scan(&x, &y); err != nil {
		context.WithField("When", "Finding tile").Error(err.Error())
	}
	return x, y, continentX, continentY, err
}
