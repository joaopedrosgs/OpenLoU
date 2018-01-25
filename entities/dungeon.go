package entities

import (
	"time"
)

type Dungeon struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	TileNode
	Level    int
	Progress int
}
