package models

import (
	"github.com/jackc/pgx"
	"time"
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

func GetUpgradesFromCity(db pgx.Conn, tile Coord, continent Coord) []*Upgrade {

}

func GetAllUpgrades(db pgx.Conn) []*Upgrade {

}
