package database

import (
	"errors"
	"github.com/joaopedrosgs/OpenLoU/entities"
)

func CreateConstruction(cityID, x, y, constructionType, level uint) error {
	numberOfUpgradesCity := 0
	err := db.QueryRow("(SELECT COUNT(DISTINCT index) FROM upgrades WHERE city_id = $1)", cityID).Scan(&numberOfUpgradesCity)
	if err != nil {
		return errors.New("failed to get number of upgrades of this city:" + err.Error())
	}
	if numberOfUpgradesCity >= 10 {
		return errors.New("construction queue is full")
	}
	_, err = db.Exec("INSERT into constructions (city_id, x, y, type, level) VALUES($1,$2,$3,$4,$5)", cityID, x, y, constructionType, level)
	return err
}
func GetConstruction(cityID, x, y uint) (*entities.Construction, error) {
	construction := entities.Construction{}
	err := db.QueryRow("Select * from constructions WHERE city_id = $1 AND x = $2 AND y = $3", cityID, x, y).Scan(
		&construction.ID,
		&construction.CreatedAt,
		&construction.UpdatedAt,
		&construction.CityID,
		&construction.Level,
		&construction.X,
		&construction.Y,
		&construction.Type,
		&construction.Production,
		&construction.Modifier,
		&construction.NeedRefresh)
	if err != nil {
		return nil, err
	}
	return &construction, err
}
func GetAllConstruction(cityID uint) (*[]entities.Construction, error) {
	constructions := make([]entities.Construction, 0, 100)
	rows, err := db.Query("Select * from constructions WHERE city_id = $1", cityID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		construction := entities.Construction{}
		rows.Scan(
			&construction.ID,
			&construction.CreatedAt,
			&construction.UpdatedAt,
			&construction.CityID,
			&construction.Level,
			&construction.X,
			&construction.Y,
			&construction.Type,
			&construction.Production,
			&construction.Modifier,
			&construction.NeedRefresh)
		constructions = append(constructions, construction)
	}

	return &constructions, err
}
func CreateUpgrade(cityId, x, y uint) error {
	constructionID, level := 0, 0
	err := db.QueryRow("SELECT id, level FROM constructions WHERE city_id = $1 AND x = $2 AND y = $3 limit 1", cityId, x, y).Scan(&constructionID, &level)
	if err != nil {
		return errors.New("Failed to create upgrade:" + err.Error())
	}
	if level >= 10 {
		return errors.New("already at max level")
	}
	numberOfUpgrades := 0

	err = db.QueryRow("(SELECT COUNT(DISTINCT index) FROM upgrades WHERE construction_id = $1)", constructionID).Scan(&numberOfUpgrades)
	if err != nil {
		return errors.New("failed to get number of upgrades of this construction:" + err.Error())
	}
	if level+numberOfUpgrades >= 10 {
		return errors.New("already at max level")
	}
	numberOfUpgradesCity := 0
	err = db.QueryRow("(SELECT COUNT(DISTINCT index) FROM upgrades WHERE city_id = $1)", cityId).Scan(&numberOfUpgradesCity)
	if err != nil {
		return errors.New("failed to get number of upgrades of this city:" + err.Error())
	}
	if numberOfUpgradesCity >= 10 {
		return errors.New("construction queue is full")
	}
	_, err = db.Exec("INSERT into upgrades(city_id, construction_id, index, duration) VALUES ($1, $2, $3, $4)", cityId, constructionID, numberOfUpgrades+1, 10)
	if err != nil {
		return errors.New("failed to create upgrade:" + err.Error())
	}
	return nil
}
func GetUpgrades() (*[]entities.Upgrade, error) {
	upgrades := make([]entities.Upgrade, 0, 100)
	rows, err := db.Query("SELECT * FROM upgrades WHERE start+(duration * interval '1 second') > now() AND index = 1 LIMIT 100")
	defer rows.Close()
	if err != nil {
		return nil, errors.New("Failed to get upgrades:" + err.Error())
	}
	for rows.Next() {
		upgrade := entities.Upgrade{}
		err := rows.Scan(&upgrade.ID, &upgrade.CreatedAt, &upgrade.UpdatedAt,
			&upgrade.ConstructionID, &upgrade.CityID, &upgrade.Index, &upgrade.Duration, &upgrade.Start)
		if err != nil {
			return nil, errors.New("Failed to get upgrades:" + err.Error())
		}

		upgrades = append(upgrades, upgrade)
	}
	return &upgrades, nil
}
func CompleteUpgrade(upgrade entities.Upgrade) error {
	_, err := db.Exec("DELETE FROM upgrades WHERE id = $1", upgrade.ID)
	if err != nil {
		return errors.New("Failed to delete upgrade:" + err.Error())
	}
	_, err = db.Exec("UPDATE constructions SET level = level + 1 WHERE id = $1", upgrade.ConstructionID)
	if err != nil {
		return errors.New("Failed to set level after upgrade:" + err.Error())
	}
	_, err = db.Exec("UPDATE upgrades SET index = index - 1 WHERE city_id = $1", upgrade.CityID)
	if err != nil {
		return errors.New("Failed to set index after upgrade:" + err.Error())
	}
	_, err = db.Exec("UPDATE upgrades SET start = now() WHERE city_id = $1 AND index = 1", upgrade.CityID)
	if err != nil {
		return errors.New("Failed to start next upgrade after upgrade:" + err.Error())
	}

	return err
}
