package models

import (
	"time"
)

type Terrain struct {
	Tile           Coord
	ContinentCoord Coord
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Type           uint
}
