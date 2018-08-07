package storage

import (
	"math"

	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
)

// GetAllDungeons recupera todas as dungeons do jogo do banco de dados
func GetAllDungeons(db *pgx.Conn) ([]*models.Dungeon, error) {
	dungeons := make([]*models.Dungeon, 0, 2000)
	rows, err := db.Query("SELECT * FROM DUNGEONS")
	if err != nil {
		logger.Error("failed to query all dungeons: ", err.Error)
		return nil, err
	}
	for rows.Next() {
		dungeon := &models.Dungeon{}
		err = rows.Scan(
			&dungeon.Tile.X,
			&dungeon.Tile.Y,
			&dungeon.ContinentTile.X,
			&dungeon.ContinentTile.Y,
			&dungeon.Type,
			&dungeon.Level,
			&dungeon.Progress)
		if err != nil {
			logger.Error("failed to scan all dungeons: ", err.Error)
			return nil, err
		}
		dungeons = append(dungeons, dungeon)
	}
	return dungeons, nil

}

// GetAllDungeonsInRange recupera todas as dungeons do jogo em um determinado raio
func GetAllDungeonsInRange(db *pgx.Conn, tile models.Coord, continentTile models.Coord, rang int) ([]*models.Dungeon, error) {
	dungeons := make([]*models.Dungeon, 0, int(math.Pow(float64(rang*2), 2)))
	rows, err := db.Query("SELECT * FROM DUNGEONS where x > $1 and x<$2 and y > $3 and y<$4", tile.X-rang, tile.X+rang, tile.Y-rang, tile.Y+rang)
	if err != nil {
		logger.Error("failed to query all dungeons in range: ", err.Error)

		return nil, err
	}
	for rows.Next() {
		dungeon := &models.Dungeon{}
		err = rows.Scan(
			&dungeon.Tile.X,
			&dungeon.Tile.Y,
			&dungeon.ContinentTile.X,
			&dungeon.ContinentTile.Y,
			&dungeon.Type,
			&dungeon.Level,
			&dungeon.Progress)
		if err != nil {
			logger.Error("failed to scan all dungeons in range: ", err.Error)
			return nil, err
		}
		dungeons = append(dungeons, dungeon)
	}
	return dungeons, nil
}

// GetDungeon recupera uma dungeon especifica do banco de dados
func GetDungeon(db *pgx.Conn, tile models.Coord, continentTile models.Coord) (*models.Dungeon, error) {
	dungeon := &models.Dungeon{}
	row := db.QueryRow("SELECT * FROM DUNGEONS where x = $1 AND y = $2 ", tile.X, tile.Y)
	err := row.Scan(
		&dungeon.Tile.X,
		&dungeon.Tile.Y,
		&dungeon.ContinentTile.X,
		&dungeon.ContinentTile.Y,
		&dungeon.Type,
		&dungeon.Level,
		&dungeon.Progress)
	if err != nil {
		logger.Error("failed to scan dungeon: ", err.Error)
		return nil, err
	}
	return dungeon, err
}
