package database

import (
	"errors"
	"github.com/joaopedrosgs/OpenLoU/entities"
)

func CreateConstruction(construction *entities.Construction) error {
	/*
		numberOfUpgradesCity := 0
		err := db.QueryRow(
			"SELECT COUNT(DISTINCT index) "+
				"FROM upgrades "+
				"WHERE city_x = $1 and city_y = $2",
			construction.CityX,
			construction.CityY).Scan(&numberOfUpgradesCity)
		if err != nil {
			return errors.New("failed to get number of upgrades of this city:" + err.Error())
		}
		if numberOfUpgradesCity >= 10 {
			return errors.New("construction queue is full")
		}
	*/
	_, err := db.Exec(
		"INSERT into constructions (city_x, city_y, x, y, type, level) "+
			"VALUES($1, $2, $3, $4, $5, $6)",
		construction.CityX,
		construction.CityY,
		construction.X,
		construction.Y,
		construction.Type,
		construction.Level)
	return err
}
func GetConstruction(cityX, cityY, x, y uint) (*entities.Construction, error) {
	construction := entities.Construction{}
	err := db.QueryRow(
		"Select x, y, city_x, city_y, level, type, production "+
			"from constructions "+
			"WHERE city_x = $1 AND city_y = $2 AND x = $3 AND y = $4", cityX, cityY, x, y).Scan(
		&construction.X,
		&construction.Y,
		&construction.CityX,
		&construction.CityY,
		&construction.Level,
		&construction.Type,
		&construction.Production)
	return &construction, err
}
func GetAllConstruction(cityX, cityY uint) (*[]entities.Construction, error) {
	constructions := make([]entities.Construction, 0, 100)
	rows, err := db.Query("Select x, y, city_x, city_y, level, type, production from constructions WHERE city_x = $1 AND city_y = $2", cityX, cityY)
	if err == nil {
		for rows.Next() {
			construction := entities.Construction{}
			err := rows.Scan(
				&construction.X,
				&construction.Y,
				&construction.CityX,
				&construction.CityY,
				&construction.Level,
				&construction.Type,
				&construction.Production)
			if err != nil {
				context.Error(err.Error())
				break
			}
			constructions = append(constructions, construction)

		}
	} else {
		context.Error(err.Error())
	}
	return &constructions, err
}
func CreateUpgrade(construction *entities.Construction) error {
	err := db.QueryRow(""+
		"SELECT level FROM constructions "+
		"WHERE city_x = $1 AND city_y = $2 AND x = $3 AND y = $4 "+
		"limit 1",
		construction.CityX,
		construction.CityY,
		construction.X,
		construction.Y).Scan(&construction.Level)
	if err != nil {
		return errors.New("Failed to create upgrade:" + err.Error())
	}
	if construction.Level > 9 {
		return errors.New("already at max level")
	}
	var numberOfUpgrades uint

	err = db.QueryRow(
		"SELECT COUNT(DISTINCT index) "+
			"FROM upgrades "+
			"WHERE city_x = $1 AND city_y = $2 AND x = $3 AND y = $4",
		construction.CityX,
		construction.CityY,
		construction.X,
		construction.Y).Scan(&numberOfUpgrades)
	if err != nil {
		return errors.New("failed to get number of upgrades of this construction:" + err.Error())
	}
	if construction.Level+numberOfUpgrades > 9 {
		return errors.New("already at max level")
	}
	numberOfUpgradesCity := 0
	err = db.QueryRow(""+
		"SELECT COUNT(DISTINCT index) "+
		"FROM upgrades "+
		"WHERE city_x = $1 AND city_y = $2", construction.CityX, construction.CityY).Scan(&numberOfUpgradesCity)
	if err != nil {
		return errors.New("failed to get number of upgrades of this city:" + err.Error())
	}
	if numberOfUpgradesCity > 9 {
		return errors.New("construction queue is full")
	}
	_, err = db.Exec(
		"INSERT into upgrades(city_x, city_y, x, y, index, duration)"+
			" VALUES ($1, $2, $3, $4, $5, $6)",
		construction.CityX,
		construction.CityY,
		construction.Y,
		construction.Y,
		numberOfUpgrades+1,
		10)
	if err != nil {
		return errors.New("failed to create upgrade:" + err.Error())
	}
	return nil
}
func GetUpgrades() (*[]entities.Upgrade, error) {
	upgrades := make([]entities.Upgrade, 0, 100)
	rows, err := db.Query(
		"SELECT x, y, city_x, city_y " +
			"FROM upgrades " +
			"WHERE start+(duration * interval '1 second') > now() AND index = 1 " +
			"LIMIT 100")
	defer rows.Close()
	if err != nil {
		return nil, errors.New("Failed to get upgrades:" + err.Error())
	}
	for rows.Next() {
		upgrade := entities.Upgrade{}
		err := rows.Scan(
			&upgrade.X,
			&upgrade.Y,
			&upgrade.CityX,
			&upgrade.CityY)
		if err != nil {
			return nil, errors.New("Failed to get upgrades:" + err.Error())
		}

		upgrades = append(upgrades, upgrade)
	}
	return &upgrades, nil
}

func GetUpgradesFromCity(x, y uint) (*[]entities.Upgrade, error) {
	upgrades := make([]entities.Upgrade, 0, 100)
	rows, err := db.Query(
		"SELECT x, y, city_x, city_y "+
			"FROM upgrades "+
			"WHERE city_x = $1 and city_y = $2", x, y)
	defer rows.Close()
	if err != nil {
		return nil, errors.New("Failed to get upgrades:" + err.Error())
	}
	for rows.Next() {
		upgrade := entities.Upgrade{}
		err := rows.Scan(
			&upgrade.X,
			&upgrade.Y,
			&upgrade.CityX,
			&upgrade.CityY)
		if err != nil {
			return nil, errors.New("Failed to get upgrades:" + err.Error())
		}

		upgrades = append(upgrades, upgrade)
	}
	return &upgrades, nil
}
func CompleteUpgrade(upgrade entities.Upgrade) error {
	_, err := db.Exec(
		"DELETE FROM upgrades "+
			"WHERE x = $1 AND y = $2 AND city_x = $3 AND city_y = $4", upgrade.X, upgrade.Y, upgrade.CityX, upgrade.CityY)
	if err != nil {
		return errors.New("Failed to delete upgrade:" + err.Error())
	}
	_, err = db.Exec(
		"UPDATE constructions "+
			"SET level = level + 1 "+
			"WHERE x = $1 AND y = $2 AND city_x = $3 AND city_y = $4", upgrade.X, upgrade.Y, upgrade.CityX, upgrade.CityY)
	if err != nil {
		return errors.New("Failed to set level after upgrade:" + err.Error())
	}
	_, err = db.Exec(
		"UPDATE upgrades "+
			"SET index = index - 1 "+
			"WHERE city_x = $1 AND city_y = $2", upgrade.CityX, upgrade.CityY)
	if err != nil {
		return errors.New("Failed to set index after upgrade:" + err.Error())
	}
	_, err = db.Exec(
		"UPDATE upgrades "+
			"SET start = now() "+
			"WHERE city_x = $1 AND city_y = $2 AND index = 1", upgrade.CityX, upgrade.CityY)
	if err != nil {
		return errors.New("Failed to start next upgrade after upgrade:" + err.Error())
	}

	return err
}
