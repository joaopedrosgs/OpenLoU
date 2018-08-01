package storage

import (
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
)

func GetAllTiles(db pgx.Conn) []*models.Terrain {

}
func GetAllTilesInRange(db pgx.Conn, tile models.Coord, continent models.Coord, rang uint) []*models.Terrain {

}

func GetTile(db pgx.Conn, tile models.Coord, continent models.Coord) *models.Terrain {

}
