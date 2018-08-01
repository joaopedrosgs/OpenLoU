package database

import (
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/models"
	"math"
	"math/rand"
	"time"
)

func createNewContinents() error {
	numberOfContinents := uint(math.Sqrt(float64(configuration.GetSingleton().Parameters.General.WorldSize)))
	continentSize := configuration.GetSingleton().Parameters.General.ContinentSize
	groundType := [4]string{"land", "stone", "iron", "forest"}

	_, err := db.Exec("DELETE from tiles")
	if err != nil {
		context.WithField("When", "Deleting tiles").Error(err.Error())
		return err
	}
	_, err = db.Exec("DELETE from continents")
	if err != nil {
		context.WithField("When", "Deleting continents").Error(err.Error())
		return err
	}

	var x, y uint
	for x = 0; x < numberOfContinents; x++ {
		for y = 0; y < numberOfContinents; y++ {
			_, err = db.Exec("INSERT into continents (is_active ,size, x, y) values(true, $1, $2, $3)", continentSize, x, y)
			if err != nil {
				context.WithField("When", "Inserting new continent").Error(err.Error())
				return err
			}
			rand.Seed(time.Now().UTC().UnixNano())
			var tileX, tileY uint
			for tileX = x * 100; tileX < (x+1)*100; tileX++ {
				for tileY = y * 100; tileY < (y+1)*100; tileY++ {
					go func(x, y, tileX, tileY uint) {
						_, err := db.Exec(
							"INSERT into tiles (continent_x, continent_y, x, y, occupied_by) "+
								"values ($1, $2, $3, $4, $5)", x, y, tileX, tileY, groundType[rand.Intn(3)])
						if err != nil {
							context.WithField("When", "Inserting new tile").Error(err.Error())
						}
					}(x, y, tileX, tileY)

				}
			}
		}
	}
	context.Info("Continents created!")
	return nil
}
func setTile(tile models.TileNode) error {
	_, err := db.Exec("UPDATE tiles SET occupied_by = $1 WHERE continent_x = $2 AND continent_y = $3 AND x = $4 AND y = $5", tile.Type, tile.ContinentX, tile.ContinentY, tile.X, tile.Y)
	return err
}

func checkContinentSpace() (bool, error) {
	available := true
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM continents WHERE is_active AND number_of_cities < 200)").Scan(&available)
	return available, err
}
func startingNewContinent() error {
	context.Info("Setting up new continent")
	_, err := db.Exec("UPDATE continents SET is_active = FALSE WHERE is_active = TRUE ")
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE continents " +
		"SET is_active = TRUE " +
		"FROM (SELECT x,y from continents where number_of_cities < 200 ORDER by x,y asc limit 1) SelectedContinent " +
		"WHERE continents.x = SelectedContinent.x AND" +
		"continents.y = SelectedContinent.y")
	if err != nil {
		return err
	}
	return nil
}
