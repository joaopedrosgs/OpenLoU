package entities

import "github.com/jinzhu/gorm"

type Continent struct {
	gorm.Model
	IsActive    bool
	Size        int
	CitiesLimit int
	X           int `gorm:"index:continent_idx_x_y"`
	Y           int `gorm:"index:continent_idx_x_y"`
}
