package entities

import "github.com/jinzhu/gorm"

type Continent struct {
	gorm.Model
	IsActive    bool
	Size        int
	CitiesLimit int
	X           int
	Y           int
	Tiles       []TileNode `gorm:"ForeignKey:ContinentID; "`
}
