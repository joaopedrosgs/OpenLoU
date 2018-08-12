package storage

import (
	"context"
	"math"
	"math/rand"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"github.com/joaopedrosgs/OpenLoU/configuration"
	"github.com/joaopedrosgs/OpenLoU/models"
)

// GetAllContinents recupera todos continentes do jogo
func GetAllContinents(db *pgx.Conn) ([]*models.Continent, error) {
	continents := make([]*models.Continent, 0, 36)
	rows, err := db.Query("SELECT * FROM continents")
	if err != nil {
		logger.Error("failed to get all continents: ", err.Error)
		return nil, err
	}
	for rows.Next() {
		continent := &models.Continent{}
		err := rows.Scan(
			&continent.Tile.X,
			&continent.Tile.Y,
			&continent.CreatedAt,
			&continent.UpdatedAt,
			&continent.IsActive,
			&continent.Size,
			&continent.NumberOfCities,
			&continent.CitiesLimit)
		if err != nil {
			logger.Error("failed to scan all continents: ", err.Error())
			return nil, err
		}
		continents = append(continents, continent)
	}
	return continents, nil
}

// GetContinent recupera um determinado continent do jogo
func GetContinent(db *pgx.Conn, tile models.Coord) (*models.Continent, error) {
	continent := &models.Continent{}
	row := db.QueryRow("SELECT * FROM continents WHERE x=$1 AND y=$2", tile.X, tile.Y)
	err := row.Scan(
		&continent.Tile.X,
		&continent.Tile.Y,
		&continent.CreatedAt,
		&continent.UpdatedAt,
		&continent.IsActive,
		&continent.Size,
		&continent.NumberOfCities,
		&continent.CitiesLimit)
	if err != nil {
		logger.Error("failed do scan continent: ", err.Error)
		return nil, err
	}
	return continent, nil
}

func CreateNewContinents(db *pgx.Conn) error {
	logger.Info("Populating continents and tiles")
	numberOfContinents := uint(math.Sqrt(float64(configuration.GetSingleton().Parameters.General.WorldSize)))
	continentSize := configuration.GetSingleton().Parameters.General.ContinentSize
	groundType := [4]string{"land", "stone", "iron", "forest"}

	_, err := db.Exec("DELETE from tiles")
	if err != nil {
		logger.WithField("When", "Deleting tiles").Error(err.Error())
		return err
	}
	_, err = db.Exec("DELETE from continents")
	if err != nil {
		logger.WithField("When", "Deleting continents").Error(err.Error())
		return err
	}

	var x, y uint
	start := time.Now()

	for x = 0; x < numberOfContinents; x++ {
		for y = 0; y < numberOfContinents; y++ {
			_, err = db.Exec("INSERT into continents (is_active ,size, x, y) values(true, $1, $2, $3)", continentSize, x, y)
			if err != nil {
				logger.WithField("When", "Inserting new continent").Error(err.Error())
				return err
			}
			batch := db.BeginBatch()

			rand.Seed(time.Now().UTC().UnixNano())
			var tileX, tileY uint
			for tileX = x * 100; tileX < (x+1)*100; tileX++ {

				for tileY = y * 100; tileY < (y+1)*100; tileY++ {
					batch.Queue(
						"INSERT into tiles (continent_x, continent_y, x, y, occupied_by) values ($1, $2, $3, $4, $5)",
						[]interface{}{x, y, tileX, tileY, groundType[rand.Intn(3)]},
						[]pgtype.OID{pgtype.Int2OID, pgtype.Int2OID, pgtype.Int2OID, pgtype.Int2OID, pgtype.VarcharOID},
						nil)

				}

			}
			batch.Send(context.Background(), nil)

		}
	}
	if err != nil {
		logger.Fatal("failed to create continents", err.Error())
		return err
	}
	logger.Info("Continents created!")
	logger.Info(time.Since(start))

	return nil
}
