package entities

import (
	"time"
)

type WorldResource struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	TileNode
	ResourceType uint
}
