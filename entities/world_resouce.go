package entities

import "github.com/jinzhu/gorm"

type WorldResource struct {
	gorm.Model
	TileNode
	ResourceType uint
}
