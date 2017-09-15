package gameEntities

import "time"

type cQueueNode struct {
	ContructionType uint
	Completion      time.Time
	PosX            uint8
	PosY            uint8
}
