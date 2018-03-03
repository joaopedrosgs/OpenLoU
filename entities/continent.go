package entities

import (
	"time"
)

type Continent struct {
	X              uint
	Y              uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsActive       bool
	Size           uint
	NumberOfCities uint
	CitiesLimit    uint
}
