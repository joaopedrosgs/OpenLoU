package entities

import (
	"time"
)

type Continent struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsActive       bool
	Size           uint
	NumberOfCities uint
	CitiesLimit    uint
	X              uint
	Y              uint
}
