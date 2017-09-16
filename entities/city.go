package entities

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type City struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	OwnerID     uint
	CityName    string
	Points      uint
	IsCastle    bool
	IsPalace    bool
	PosX        uint
	PosY        uint
	data        cityData
	lastUpdated time.Time
}

func (c *City) GetData() *cityData {
	return &c.data
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
	ID        uint
	FromID    uint
	ToID      uint
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
	TotalRes      [4]uint
	ActualRes     [4]uint
	Transports    []*Transport
	Troops []struct {
		AtBase []struct {
			Type     uint8
			Quantity uint
		}
		Moving []struct {
			Type     uint8
			Quantity uint
		}
	}
	Carts struct {
		AtBase uint
		Moving uint
	}
	Ships struct {
		AtBase uint
		Moving uint
	}
}

func (city *City) InRange(posX uint, posY uint, rang uint) bool {
	if city.PosX >= posX-rang &&
		city.PosX <= posX+rang &&
		city.PosY <= posY+rang &&
		city.PosY >= posY-rang {
		return true
	}
	return false
}

func (c *City) BuildEnqueue(_due Due) bool {
	if len(c.GetData().Queue.Constructions) > 10 {
		return false
	} else if len(c.GetData().Queue.Constructions) == 0 {
		_due.Start = time.Now()
	}
	c.data.Queue.Constructions = append(c.data.Queue.Constructions, _due)
	return true
}

func (c *City) TroopEnqueue(_due Due) error {

	for res := 0; res < 4; res++ {
		if RegisteredTroops[uint(_due.Type)].Cost[res]*uint(_due.Value) > c.data.ActualRes[res] {
			return errors.New("Insufficient resources!")
		}
	}

	if len(c.GetData().Queue.Military) > 5 {
		return errors.New("Limite de atividades militares excedido")
	} else if len(c.GetData().Queue.Military) == 0 {
		_due.Start = time.Now()
	}
	c.data.Queue.Military = append(c.data.Queue.Military, _due)
	return nil
}

func (c *City) TransportEnqueue(_transport Transport) ([]*Transport, error) {

	var total uint
	for _, v := range _transport.Resources {
		total += v
	}

	if _transport.Water {
		shipsNeeded := total % 10000
		if c.data.Ships.AtBase > shipsNeeded {
			c.data.Ships.AtBase -= shipsNeeded
			c.data.Ships.Moving += shipsNeeded
		} else {
			return nil, errors.New("Sem navios suficientes")
		}
	} else {
		cartsNeeded := total % 1000
		if c.data.Carts.AtBase > cartsNeeded {
			c.data.Carts.AtBase -= cartsNeeded
			c.data.Carts.Moving += cartsNeeded
		} else {
			return nil, errors.New("Sem carroças suficientes")
		}

	}

	c.data.Transports = append(c.data.Transports, &_transport)

	return c.data.Transports, nil

}
