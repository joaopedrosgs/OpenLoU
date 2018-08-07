package models

import (
	"time"
)

type Dungeon struct {
	Tile          Coord
	ContinentTile Coord
	Type          int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Level         int
	Progress      int
}
