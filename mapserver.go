package main

import (
	"LordOfUltima/city"
	"LordOfUltima/database"
	"LordOfUltima/dungeon"

	_ "time"
)

type mapServer struct {
	Cities    []city.City
	Dungeons  []dungeon.Dungeon
	Resources []Resource
	Requests  chan Request
	Answers   chan Answer
}

func (m *mapServer) LoadAndStart() {
	m.Requests = make(chan Request)
	m.Answers = make(chan Answer)
	db := database.Session.DB("lordofultima")
	db.C("cities").Find(nil).All(&m.Cities)
	go m.RequestProcessor()
	m.Requests <- Request{0, 2, 3}
}

func (m *mapServer) getCities(posX uint, posY uint, rang uint) []city.City {
	var listReturn []city.City
	for i := range m.Cities {
		if m.Cities[i].InRange(posX, posY, rang) {
			listReturn = append(listReturn, m.Cities[i])
		}
	}
	return listReturn

}

func (m *mapServer) RequestProcessor() {

	request := <-m.Requests
	switch request.Code {
	case 0:
		println("caso 0")
		break
	}
}
