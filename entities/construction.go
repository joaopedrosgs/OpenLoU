package entities

import (
	"time"
)

type Construction struct {
	X           uint
	Y           uint
	CityX       uint
	CityY       uint
	Type        uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Level       uint
	Production  int
	Modifier    int
	NeedRefresh bool
}

type Upgrade struct {
	X                uint
	Y                uint
	CityX            uint
	CityY            uint
	ConstructionType uint
	Downgrade        bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Index            uint
	Duration         uint
	Start            time.Time
}
