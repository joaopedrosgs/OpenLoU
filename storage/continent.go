package storage

import (
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
)

func GetAllContinents(db *pgx.Conn) []*models.Continent {
	continents := make([]*models.Continent, 0, 36)
	return continents
}

func GetContinent(db *pgx.Conn, tile models.Coord) *models.Continent {
	continent := &models.Continent{}
	return continent
}
