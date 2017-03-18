package city

import (
	"LordOfUltima/constructions"
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

type cityData struct {
	constructions constructions.ConstructionType
	comentary     string
	production    [4]uint
	total         [4]uint
	actual        [4]int
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
