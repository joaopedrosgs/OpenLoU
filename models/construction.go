package models

import (
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

type Construction struct {
	Tile        Coord
	CityTile    Coord
	Type        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Level       int
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

func (c *Construction) GetCoord() Coord {
	return Coord{X: c.CityTile.X<<2 + c.Tile.X, Y: c.CityTile.Y<<2 + c.Tile.Y}
}

func (c *Construction) UpgradeConstruction(db *pgx.Batch) {

	db.Queue("UPDATE constructions SET level = level+1, WHERE x=$1 AND y=$2 AND city_x = $3 AND city_y=$4",
		[]interface{}{c.Tile.X,
			c.Tile.Y,
			c.CityTile.X,
			c.CityTile.Y},
		[]pgtype.OID{pgtype.Int2OID, pgtype.Int2OID, pgtype.Int2OID, pgtype.Int2OID},
		nil)

}

func (c *Construction) MoveConstruction(db *pgx.Conn, newTile Coord) error {

	_, err := db.Exec(
		"UPDATE constructions SET x = $1, y =  $2 WHERE x = $3 AND y = $4 AND city_x = $5 AND city_y = $6",
		newTile.X,
		newTile.Y,
		c.Tile.X,
		c.Tile.Y,
		c.CityTile.X,
		c.CityTile.Y)

	return err
}
