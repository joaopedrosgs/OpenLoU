package constructions

import "time"

type QueueNode struct {
	ContructionType uint
	Completion      time.Time
	PosX            uint8
	PosY            uint8
}
