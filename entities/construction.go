package entities

import "time"

type Construction struct {
	ID          int
	CityID      int
	Level       int
	X           int
	Y           int
	Type        int
	Production  int
	Modifier    int
	needRefresh bool
}

type Upgrade struct {
	ConstructionID int
	ToLevel        int
	Completion     time.Time
}
