package entities

import (
	"time"
)

type Dungeon struct {
	TileNode
	CreatedAt time.Time
	UpdatedAt time.Time
	Level     int
	Progress  int
}
