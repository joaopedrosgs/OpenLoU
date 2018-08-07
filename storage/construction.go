package storage

import (
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
)

// GetConstruction retorna a construção de uma determinada cidade
func GetConstruction(db *pgx.Conn, tile models.Coord, cityTile models.Coord) (*models.Construction, error) {
	construction := &models.Construction{}
	err := db.QueryRow(
		"SELECT x, y, city_x, city_y, level, type, production "+
			"FROM constructions "+
			"WHERE city_x = $1 AND city_y = $2 AND x = $3 AND y = $4", cityTile.X, cityTile.Y, tile.X, tile.Y).Scan(
		&construction.Tile.X,
		&construction.Tile.Y,
		&construction.CityTile.X,
		&construction.CityTile.Y,
		&construction.Level,
		&construction.Type,
		&construction.Production)
	return construction, err
}

// GetAllConstructions retorna todas as construções de uma determinada cidade
func GetAllConstructions(db *pgx.Conn, cityTile models.Coord) ([]*models.Construction, error) {
	constructions := make([]*models.Construction, 0, 100)
	rows, err := db.Query("SELECT x, y, city_x, city_y, level, type, production FROM constructions WHERE city_x = $1 AND city_y = $2", cityTile.X, cityTile.Y)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	for rows.Next() {
		construction := &models.Construction{}
		err := rows.Scan(
			&construction.Tile.X,
			&construction.Tile.Y,
			&construction.CityTile.X,
			&construction.CityTile.Y,
			&construction.Level,
			&construction.Type,
			&construction.Production)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		constructions = append(constructions, construction)

	}

	return constructions, err
}

// CreateConstruction cria uma construção em uma determinada cidade
func CreateConstruction(db *pgx.Conn, construction *models.Construction) error {

	_, err := db.Exec(
		"INSERT into constructions (city_x, city_y, x, y, type, level) "+
			"VALUES($1, $2, $3, $4, $5, $6)",
		construction.CityTile.X,
		construction.CityTile.Y,
		construction.Tile.X,
		construction.Tile.Y,
		construction.Type,
		construction.Level)
	return err
}
