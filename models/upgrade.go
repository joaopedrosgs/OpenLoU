package models

import (
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

type Upgrade struct {
	Tile             Coord
	CityTile         Coord
	ConstructionType uint
	Downgrade        bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Index            uint
	Duration         uint
	Start            time.Time
}

func (upgrade *Upgrade) EnqueueCompletion(db *pgx.Batch) {
	construction := Construction{Tile: upgrade.Tile, CityTile: upgrade.CityTile}

	construction.UpgradeConstruction(db)
	db.Queue("DELETE FROM upgrades WHERE x=$1 and y=$2 and city_x=$3 and city_y=$4",
		[]interface{}{upgrade.Tile.X,
			upgrade.Tile.Y,
			upgrade.CityTile.X,
			upgrade.CityTile.Y},
		[]pgtype.OID{pgtype.Int2OID, pgtype.Int2OID, pgtype.Int2OID, pgtype.Int2OID},
		nil)
	return
}
