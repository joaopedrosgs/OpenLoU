package gameEntities

import "time"

type mQueueNode struct {
	troopType  uint
	completion time.Time
	quantity   uint
}
