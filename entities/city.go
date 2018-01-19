package entities

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type City struct {
	ID uint `gorm:"primary_key"`
	TileNode
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
	UserID          uint `gorm:"index"`
	Name            string
	Points          int
	Constructions   []Construction `gorm:"ForeignKey:CityID"`
	Upgrades        []Upgrade      `gorm:"ForeignKey:CityID"`
	WoodProduction  int
	StoneProduction int
	IronProduction  int
	FoodProduction  int
	GoldProduction  int
	WoodStored      int
	StoneStored     int
	FoodStored      int
	IronStored      int
	WoodLimit       int
	StoneLimit      int
	IronLimit       int
	FoodLimit       int
}

func (c *City) GetType() int {
	return 1
}
