package storage

import (
	"github.com/jackc/pgx"
	"github.com/joaopedrosgs/OpenLoU/models"
)

// GetUpgradesFromCity recupera todos os upgrades de uma determinada cidade
func GetUpgradesFromCity(db *pgx.Conn, cityTile models.Coord) ([]*models.Upgrade, error) {
	upgrades := make([]*models.Upgrade, 0, 10)
	rows, err := db.Query(
		"SELECT * FROM upgrades WHERE city_x = $1 AND city_y = $2",
		cityTile.X,
		cityTile.Y)
	if err != nil {
		logger.Error("Failed to query upgrades from city:", err.Error)
		return nil, err
	}
	for rows.Next() {
		upgrade := &models.Upgrade{}
		err = rows.Scan(
			&upgrade.Tile.X,
			&upgrade.Tile.Y,
			&upgrade.CityTile.X,
			&upgrade.CityTile.Y,
			&upgrade.CreatedAt,
			&upgrade.Index,
			&upgrade.Duration,
			&upgrade.Start)
		if err != nil {
			logger.Error("Failed to scan query upgrades from city:", err.Error)
			return nil, err
		}
		upgrades = append(upgrades, upgrade)
	}
	return upgrades, nil

}

// GetAllUpgrades recupera todos os upgrades de todas as cidades do jogo
func GetAllUpgrades(db *pgx.Conn) ([]*models.Upgrade, error) {
	upgrades := make([]*models.Upgrade, 0, 1000)
	rows, err := db.Query(
		"SELECT * FROM upgrades")
	if err != nil {
		logger.Error("Failed to query upgrades:", err.Error)
		return nil, err
	}
	for rows.Next() {
		upgrade := &models.Upgrade{}
		err = rows.Scan(
			&upgrade.Tile.X,
			&upgrade.Tile.Y,
			&upgrade.CityTile.X,
			&upgrade.CityTile.Y,
			&upgrade.CreatedAt,
			&upgrade.Index,
			&upgrade.Duration,
			&upgrade.Start)
		if err != nil {
			logger.Error("Failed to scan query upgrades:", err.Error)
			return nil, err
		}
		upgrades = append(upgrades, upgrade)
	}
	return upgrades, nil
}

func CreateUpgrade(db *pgx.Conn, upgrade *models.Upgrade) error {
	_, err := db.Exec("INSERT into upgrades(x, y, city_x, city_y, duration) values($1, $2, $3, $4, $5)",
		upgrade.Tile.X,
		upgrade.Tile.Y,
		upgrade.CityTile.X,
		upgrade.CityTile.Y,
		10)
	if err != nil {
		logger.Error("Failed to create upgrade: ", err.Error)
		return err
	}
	return nil
}
