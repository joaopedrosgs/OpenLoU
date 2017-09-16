package entities

import "time"

type ConstructionQueue struct {
	LastUpdate time.Time
	Items      []ConstructionQueueNode
}

type ConstructionQueueNode struct {
	ContructionType uint
	Completion      time.Time
	PosX            uint8
	PosY            uint8
}
