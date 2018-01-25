package entities

import (
	"time"
)

type City struct {
	ID uint
	TileNode
	createdAt       time.Time
	updatedAt       time.Time
	userID          uint
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

func (c *City) GetType() int {
	return 1
}
