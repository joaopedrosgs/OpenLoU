package entities

import "time"

type MilitaryQueue struct {
	LastUpdate time.Time
	Items      []MilitaryQueueNode
}

type MilitaryQueueNode struct {
	troopType  uint
	completion time.Time
	quantity   uint
}
