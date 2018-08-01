package storage

import (
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
	"math"
)

func GetAllDungeons(db pgx.Conn) []*models.Dungeon {
	dungeons := make([]*models.Dungeon, 0, 2000)
	return dungeons

}
func GetAllDungeonsInRange(db pgx.Conn, tile models.Coord, continentTile models.Coord, rang uint) []*models.Dungeon {
	dungeons := make([]*models.Dungeon, 0, math.Pow(float64(rang*2), 2))
	return dungeons
}
func GetDungeon(db pgx.Conn, tile models.Coord, continentTile models.Coord) *models.Dungeon {
	dungeon := &models.Dungeon{}
	return dungeon
}
