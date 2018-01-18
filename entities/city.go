package entities

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type City struct {
	gorm.Model
	UserID          int `gorm:"index"`
	Name            string
	ContinentID     int
	Type            int
	Points          int
	X               int            `gorm:"index:idx_x_y"`
	Y               int            `gorm:"index:idx_x_y"`
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

func (c *City) ToJson() ([]byte, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return bytes, nil
	}
	return nil, err
}

func (c *City) GetType() int {
	return 1
}

type Due struct {
	Type     uint8 // troop id or construction id
	Value    uint8 // quantity or level
	Start    time.Time
	Duration time.Duration
}
type Queue struct {
	Military      []Due
	Constructions []Due
}

type Transport struct {
	ID        int
	FromID    int
	ToID      int
	Water     bool
	Resources [5]uint
	Depart    time.Time
	Duration  time.Duration
}
