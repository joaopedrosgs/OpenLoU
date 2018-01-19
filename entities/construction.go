package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Construction struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	CityID      uint `gorm:"index"`
	Level       int
	X           uint `gorm:"index:idx_x_y"`
	Y           uint `gorm:"index:idx_x_y"`
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
