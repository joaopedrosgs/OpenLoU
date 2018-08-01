package models

import (
	"github.com/jackc/pgx"
	"time"
)

type Construction struct {
	Tile        Coord
	CityTile    Coord
	Type        uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Level       uint
	Production  int
	Modifier    float32
	NeedRefresh bool
}

type ConstructionType struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Bonus []struct {
		Name  string `json:"name"`
		Value []int  `json:"value"`
	} `json:"bonus"`
	Adjascent []struct {
		Builds []string `json:"builds"`
		Bonus  []int    `json:"bonus"`
	} `json:"adjascent"`
	ResourceCost [][]int `json:"resourceCost"`
	Score        []int   `json:"score"`
}

func (c *Construction) CanUpgrade() bool {
	return c.Level < 10
}

func (c *Construction) MoveTo(db *pgx.Conn, tile Coord) bool {
	return false
}

func (c *Construction) GetCoord() Coord {
	return Coord{X: c.CityTile.X<<2 + c.Tile.X, Y: c.CityTile.Y<<2 + c.Tile.Y}
}
