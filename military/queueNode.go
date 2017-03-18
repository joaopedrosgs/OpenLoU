package military

import "time"

type QueueNode struct {
	troopType  uint
	completion time.Time
	quantity   uint
}
