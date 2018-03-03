package entities

import (
	"time"
)

type City struct {
	TileNode
	createdAt       time.Time
	updatedAt       time.Time
	UserName        string
	Name            string
	Points          int
	woodProduction  int
	stoneProduction int
	ironProduction  int
	foodProduction  int
	goldProduction  int
	woodStored      int
	stoneStored     int
	foodStored      int
	ironStored      int
	woodLimit       int
	stoneLimit      int
	ironLimit       int
	foodLimit       int
}

func (c *City) GetType() string {
	return "city"
}
