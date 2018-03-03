package entities

import (
	"time"
)

type WorldResource struct {
	TileNode
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ResourceType uint
}
