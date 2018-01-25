package entities

import (
	"time"
)

type Construction struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CityID      uint
	Level       uint
	X           uint
	Y           uint
	Type        uint
	Production  int
	Modifier    int
	NeedRefresh bool
}

type Upgrade struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ConstructionID uint
	CityID         uint
	Index          uint
	Duration       uint
	Start          time.Time
}
