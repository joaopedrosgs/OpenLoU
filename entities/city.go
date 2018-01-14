package entities

import (
	"time"
)

type City struct {
	UserName    string
	Name        string
	ContinentID int
	Type        int
	Points      int
	X           int
	Y           int
	production  [5]int
	stored      [4]int
	limit       [4]int
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

type cityData struct {
	Constructions ConstructionType
	Comentary     string
	Queue         Queue
	Production    [4]int
	TotalRes      [4]int
	ActualRes     [4]int
	Transports    []*Transport
	Troops        []struct {
		AtBase []struct {
			Type     int
			Quantity int
		}
		Moving []struct {
			Type     int
			Quantity int
		}
	}
	Carts struct {
		AtBase int
		Moving int
	}
	Ships struct {
		AtBase int
		Moving int
	}
}
