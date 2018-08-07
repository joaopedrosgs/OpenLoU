package models

import "math"

type Coord struct {
	X int
	Y int
}

func (c *Coord) DistanceBetween(other Coord) float64 {
	return math.Sqrt((float64)(math.Pow(float64(other.X-c.X), 2) + math.Pow(float64(other.Y-c.Y), 2)))
}
