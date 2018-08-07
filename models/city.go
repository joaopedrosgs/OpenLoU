package models

import (
	"time"
)

type City struct {
	Tile          Coord
	ContinentTile Coord
	Type          byte
	createdAt     time.Time
	updatedAt     time.Time
	UserName      string
	Name          string
	Points        uint
	production    [5]int
	stored        [4]int
	limit         [4]int
}

func (c City) DistanceBetween(other *City) float64 {
	return other.Tile.DistanceBetween(c.Tile)
}

func (c *City) GetCoords() Coord {
	return c.Tile
}
