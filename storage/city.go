package storage

import (
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
	"math"
)

func GetCity(db *pgx.Conn, tile models.Coord, continent models.Coord) *models.City {
	city := &models.City{}
	return city
}
func GetAllCities(db *pgx.Conn) *models.City {
	city := &models.City{}
	return city
}
func GetAllCitiesInRange(db *pgx.Conn, tile models.Coord, continent models.Coord, rang int) []*models.City {
	cities := make([]*models.City, 0, math.Pow(float64(rang*2), 2))
	return cities
}
func GetAllCitiesFromContinent(db *pgx.Conn, continent models.Coord) []*models.City {
	cities := make([]*models.City, 0, 2000)
	return cities
}

func GetAllCitiesInRangeFromContinent(db *pgx.Conn, continent models.Coord, tile models.Coord, rang int) []*models.City {
	cities := make([]*models.City, 0, math.Pow(float64(rang*2), 2))
	return cities
}
