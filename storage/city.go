package storage

import (
	"math"
	"math/rand"
	"time"

	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
)

// GetCity recupera uma cidade do banco de dados
func GetCity(db *pgx.Conn, tile models.Coord) *models.City {
	city := &models.City{}
	row := db.QueryRow(
		"SELECT x, y, continent_x, continent_y, type, name, points "+
			"FROM cities where x = $1 and y = $2", tile.X, tile.Y)
	row.Scan(&city.Tile.X, &city.Tile.Y, &city.ContinentTile.Y, &city.ContinentTile.Y, &city.Type, &city.Name, &city.Points)
	return city
}

// GetUserCities recupera todas as cidades de um determinado usuario do banco de dados
func GetUserCities(db *pgx.Conn, userName string) (*[]models.City, error) {
	var cities []models.City
	rows, err := db.Query(
		"SELECT x, y, continent_x, continent_y, type, name, points "+
			"FROM cities "+
			"WHERE user_name = $1", userName)
	logger.Error(err.Error())
	if err != nil {
		logger.Error(err.Error())
	}
	cities = make([]models.City, 0)
	for rows.Next() {
		city := models.City{UserName: userName}
		err := rows.Scan(&city.Tile.X, &city.Tile.Y, &city.ContinentTile.X, &city.ContinentTile.Y, &city.Type, &city.Name, &city.Points)
		if err != nil {
			logger.Error(err.Error())
			break
		}
		cities = append(cities, city)

	}
	rows.Close()

	return &cities, nil
}

// GetAllCities recupera todas as cidades do jogo
func GetAllCities(db *pgx.Conn) []*models.City {
	rows, err := db.Query(
		"SELECT x, y, continent_x, continent_y, type, name, points " +
			"FROM cities ")
	logger.Error(err.Error())
	if err != nil {
		logger.Error(err.Error())
	}
	cities := make([]*models.City, 0)

	for rows.Next() {
		city := &models.City{}
		err := rows.Scan(&city.Tile.X, &city.Tile.Y, &city.ContinentTile.X, &city.ContinentTile.Y, &city.Type, &city.Name, &city.Points)
		if err != nil {
			logger.Error(err.Error())
			break
		}
		cities = append(cities, city)

	}
	rows.Close()

	return cities
}

// GetAllCitiesInRange recupera todas as cidades em um determinado raio
func GetAllCitiesInRange(db *pgx.Conn, tile models.Coord, rang int) ([]*models.City, error) {
	cities := make([]*models.City, 0, int(math.Pow(float64(rang*2), 2)))
	if tile.X < rang {
		tile.X = rang
	}
	if tile.Y < rang {
		tile.Y = rang
	}
	minX := tile.X - rang
	minY := tile.Y - rang
	maxX := tile.X + rang
	maxY := tile.Y + rang
	rows, err := db.Query(
		"SELECT x, y, continent_x, continent_y, type, name, points "+
			"FROM cities "+
			"WHERE x BETWEEN $1 AND $2 AND y BETWEEN $3 AND $4",
		minX, maxX, minY, maxY)
	if err != nil {
		logger.WithField("When", "Finding cities in the region").Error(err.Error())
		return nil, err
	}
	for rows.Next() {
		city := &models.City{}
		err := rows.Scan(&city.Tile.X, &city.Tile.Y, &city.ContinentTile.X, &city.ContinentTile.Y, &city.Type, &city.Name, &city.Points)
		if err != nil {
			logger.WithField("When", "Scaning cities from the database").Error(err.Error())
			return nil, err
		}
		cities = append(cities, city)
	}

	rows.Close()
	return cities, nil
}
func populateNewCity(db *pgx.Conn, city models.City) error {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	var defaultConstruction = make([]models.Construction, 0, 21)
	defaultConstruction[0] = models.Construction{
		Tile:       models.Coord{X: 11, Y: 11},
		CityTile:   city.Tile,
		Type:       5,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Level:      1,
		Production: 300,
		Modifier:   1,
	}
	for i := 1; i < 20; i++ {
		var resourceType = random.Int() % 4
		var resourceX = random.Int() % 21
		var resourceY = random.Int() % 21
		defaultConstruction[i] = models.Construction{
			Tile:       models.Coord{X: resourceX, Y: resourceY},
			CityTile:   city.Tile,
			Type:       resourceType,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Level:      1,
			Production: 0,
			Modifier:   0,
		}
	}
	for i := 0; i < len(defaultConstruction); i++ {
		err := CreateConstruction(db, &defaultConstruction[i])
		if err != nil {
			logger.WithField("When", "Creating default constructions").Error(err.Error())
			return err
		}
	}
	return nil
}

// findNewCityLocation procura um terreno vazio para alguma finalidade
func findNewCityLocation(db *pgx.Conn) (tile models.Coord, continentTile models.Coord, err error) {

	err = db.QueryRow(
		"SELECT x,y from continents"+
			" WHERE is_active "+
			"limit 1").Scan(&continentTile.X, &continentTile.Y)
	if err != nil {
		logger.WithField("When", "Finding a continent").Error(err.Error())
		return
	}
	err = db.QueryRow(
		"SELECT x,y "+
			"FROM tiles "+
			"WHERE occupied_by = 'land' AND "+
			"continent_x = $1 AND "+
			"continent_y = $2 "+
			"ORDER by random() "+
			"limit 1", continentTile.X, continentTile.Y).Scan(&tile.X, &tile.Y)
	if err != nil {
		logger.WithField("When", "Finding tile").Error(err.Error())
	}
	return
}

// createCity cria uma cidade, populando a mesma
func createCity(db *pgx.Conn, city models.City) error {
	_, err := db.Exec(
		"INSERT INTO cities(user_name, x, y, continent_x, continent_y) "+
			"VALUES($1, $2, $3, $4, $5)",
		city.UserName, city.Tile.X, city.Tile.Y, city.ContinentTile.X, city.ContinentTile.Y)
	if err != nil {
		logger.WithField("When", "Inserting city on the database").Error(err.Error())
		return err
	}
	err = populateNewCity(db, city)
	if err != nil {
		logger.WithField("When", "Trying to populate the city").Error(err.Error())
		return err
	}
	return nil
}
