package database

import (
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"math"
	"math/rand"
	"time"
)

func createNewContinents() error {
	numberOfContinents := int(math.Sqrt(float64(configuration.GetSingleton().Parameters.General.WorldSize)))
	continentSize := configuration.GetSingleton().Parameters.General.ContinentSize
	groundType := [4]string{"land", "stone", "iron", "forest"}

	_, err := db.Exec("DELETE from continents")
	if err != nil {
		context.WithField("When", "Deleting continents").Error(err.Error())
		return err
	}

	var continentId int
	for x := 0; x < numberOfContinents; x++ {
		for y := 0; y < numberOfContinents; y++ {
			err = db.QueryRow("INSERT into continents (is_active ,size, x, y) values(true, $1, $2, $3) RETURNING id",
				continentSize, x, y).Scan(&continentId)
			if err != nil {
				context.WithField("When", "Inserting new continent").Error(err.Error())
				return err
			}
			rand.Seed(time.Now().UTC().UnixNano())
			for tileX := 0; tileX < continentSize; tileX++ {
				for tileY := 0; tileY < continentSize; tileY++ {
					go func(tileX, tileY, continentId int) {
						_, err := db.Exec("INSERT into tiles (continent_id, x, y, occupied_by) values ($1, $2, $3, $4)", continentId, tileX, tileY, groundType[rand.Intn(3)])
						if err != nil {
							context.WithField("When", "Inserting new tile").Error(err.Error())
						}
					}(tileX, tileY, continentId)

				}
			}
		}
	}
	context.Info("Continents created!")
	return nil
}
func setTile(continentID, x, y uint, to string) error {
	_, err := db.Exec("UPDATE tiles SET occupied_by = $1 WHERE continent_id = $2 AND x = $3 AND y = $4", to, continentID, x, y)
	return err
}

func checkContinentSpace() (bool, error) {
	available := true
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM continents WHERE is_active AND number_of_cities < 200)").Scan(&available)
	return available, err
}
func startingNewContinent() error {
	context.Warning("setting up new continent")
	_, err := db.Exec("UPDATE continents SET is_active = FALSE WHERE is_active = TRUE ")
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE continents SET is_active = TRUE WHERE id = (SELECT id from continents where number_of_cities < 200 ORDER by id asc limit 1)")
	if err != nil {
		return err
	}
	return nil
}
