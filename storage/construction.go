package storage

import (
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
)

func GetConstruction(db *pgx.Conn, tile models.Coord, city models.Coord) *models.Construction {
	construction := &models.Construction{}
	return construction
}
func GetAllConstrutions(db *pgx.Conn, city models.Coord) *models.Construction {
	construction := &models.Construction{}
	return construction
}
