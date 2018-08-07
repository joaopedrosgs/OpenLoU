package storage

import (
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
)

// GetAllTiles recupera todos os tiles do jogo
func GetAllTiles(db *pgx.Conn) ([]*models.Terrain, error) {
	terrains := make([]*models.Terrain, 0, 10000)
	rows, err := db.Query("SELECT * FROM tiles")
	if err != nil {
		logger.Error("failed to get all terrain: ", err.Error())
		return nil, err
	}
	for rows.Next() {
		terrain := &models.Terrain{}
		err := rows.Scan(
			&terrain.Tile.X,
			&terrain.Tile.Y,
			&terrain.ContinentTile.X,
			&terrain.ContinentTile.Y,
			&terrain.OccupiedBy)
		if err != nil {

			logger.Error("failed to scan all terrain: ", err.Error())
			return nil, err
		}
		terrains = append(terrains, terrain)

	}
	return terrains, nil
}

// GetAllTilesInRange recupera todos os tiles do jogo em um determinado raio
func GetAllTilesInRange(db *pgx.Conn, tile models.Coord, rang int) ([]*models.Terrain, error) {
	terrains := make([]*models.Terrain, 0, (rang*2)*(rang*2))
	rows, err := db.Query("SELECT * FROM tiles WHERE x >= $1 AND x<= $2 AND y>=$3 AND y<=$4",
		tile.X-rang,
		tile.X+rang,
		tile.Y-rang,
		tile.Y+rang)
	if err != nil {
		logger.Error("failed to get all terrain in range: ", err.Error())
		return nil, err
	}
	for rows.Next() {
		terrain := &models.Terrain{}
		err := rows.Scan(
			&terrain.Tile.X,
			&terrain.Tile.Y,
			&terrain.ContinentTile.X,
			&terrain.ContinentTile.Y,
			&terrain.OccupiedBy)
		if err != nil {

			logger.Error("failed to scan all terrain in range: ", err.Error())
			return nil, err
		}
		terrains = append(terrains, terrain)

	}
	return terrains, nil
}

// GetTile recupera um determinado tile do banco de dados
func GetTile(db *pgx.Conn, tile models.Coord) (*models.Terrain, error) {
	terrain := &models.Terrain{}
	rows := db.QueryRow("SELECT * FROM tiles WHERE x = $1 AND y=$2",
		tile.X,
		tile.Y)

	err := rows.Scan(
		&terrain.Tile.X,
		&terrain.Tile.Y,
		&terrain.ContinentTile.X,
		&terrain.ContinentTile.Y,
		&terrain.OccupiedBy)
	if err != nil {

		logger.Error("failed to scan terrain: ", err.Error())
		return nil, err
	}

	return terrain, nil
}
