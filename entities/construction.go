package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Construction struct {
	gorm.Model
	CityID      int `gorm:"index"`
	Level       int
	X           int `gorm:"index:idx_x_y"`
	Y           int `gorm:"index:idx_x_y"`
	Type        int
	Production  int
	Modifier    int
	NeedRefresh bool
}

type Upgrade struct {
	gorm.Model
	ConstructionID int `gorm:"index"`
	ToLevel        int
	Duration       time.Duration
	Start          time.Time
}
