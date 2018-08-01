package models

import (
	"time"
)

type Continent struct {
	Tile           Coord
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsActive       bool
	Size           uint
	NumberOfCities uint
	CitiesLimit    uint
}

