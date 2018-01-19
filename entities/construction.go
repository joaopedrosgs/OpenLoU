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
	CityID      uint `gorm:"unique_index:construction_idx_x_y_cityid"`
	Level       int
	X           uint `gorm:"unique_index:construction_idx_x_y_cityid"`
	Y           uint `gorm:"unique_index:construction_idx_x_y_cityid"`
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
